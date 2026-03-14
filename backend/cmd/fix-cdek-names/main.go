// fix-cdek-names: remove "CDEK", "СДЭК" and "арт.12345" from product names.
// Usage: go run ./cmd/fix-cdek-names
package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"uralmelal/backend/internal/config"
	"uralmelal/backend/internal/database"
	"uralmelal/backend/internal/models"
)

var artRe = regexp.MustCompile(`арт\.\s*\d+$`)

func main() {
	cfg := config.Load()
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("db: %v", err)
	}

	type result struct {
		Count int64
	}
	var r result
	updated := int64(0)

	// 1. CDEK / СДЭК — raw SQL
	db.Raw("SELECT COUNT(*) as count FROM products WHERE name LIKE ? OR name LIKE ?", "%CDEK%", "%СДЭК%").Scan(&r)
	fmt.Printf("Products with CDEK/СДЭК: %d\n", r.Count)
	res := db.Exec(`
		UPDATE products SET name = REPLACE(REPLACE(name, 'CDEK', ''), 'СДЭК', '')
		WHERE name LIKE '%CDEK%' OR name LIKE '%СДЭК%'
	`)
	if res.Error != nil {
		log.Fatalf("CDEK update: %v", res.Error)
	}
	updated += res.RowsAffected

	// 2. арт.12345 — regex (SQLite has no regex replace)
	var products []models.Product
	db.Unscoped().Where("name LIKE ?", "%арт.%").Find(&products)
	artUpdated := 0
	for i := range products {
		oldName := products[i].Name
		newName := artRe.ReplaceAllString(oldName, "")
		newName = strings.TrimSpace(newName)
		if newName != oldName {
			if err := db.Unscoped().Model(&products[i]).Update("name", newName).Error; err != nil {
				log.Printf("update product %d: %v", products[i].ID, err)
				continue
			}
			artUpdated++
		}
	}
	updated += int64(artUpdated)
	fmt.Printf("Products with арт.NNN removed: %d\n", artUpdated)

	var countCdek, countArt int64
	db.Raw("SELECT COUNT(*) FROM products WHERE name LIKE ? OR name LIKE ?", "%CDEK%", "%СДЭК%").Scan(&countCdek)
	db.Raw("SELECT COUNT(*) FROM products WHERE name LIKE ?", "%арт.%").Scan(&countArt)
	fmt.Printf("Done. Updated: %d. Remaining CDEK: %d, арт.: %d\n", updated, countCdek, countArt)
}
