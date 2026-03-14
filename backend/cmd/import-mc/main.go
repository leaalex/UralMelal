// Import MC CSV: go run ./cmd/import-mc /path/to/export_mc_ru.csv
package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"uralmelal/backend/internal/config"
	"uralmelal/backend/internal/database"
	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run ./cmd/import-mc <path-to-export_mc_ru.csv>")
		os.Exit(1)
	}
	path := os.Args[1]
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}
	defer f.Close()

	cfg := config.Load()
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("db: %v", err)
	}
	if err := db.AutoMigrate(&models.Product{}, &models.ProductAttribute{}, &models.Category{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	reader.LazyQuotes = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("read csv: %v", err)
	}
	if len(records) < 2 {
		log.Fatal("empty or header-only csv")
	}

	catRepo := repository.NewCategoryRepository(db)
	productRepo := repository.NewProductRepository(db, cfg.DBType)

	created, skipped := 0, 0
	for i := 1; i < len(records); i++ {
		row := records[i]
		if len(row) <= 2 {
			continue
		}
		name := strings.TrimSpace(at(row, 2))
		if name == "" {
			continue
		}
		sku := strings.TrimSpace(at(row, 11))
		if sku == "" {
			sku = strconv.Itoa(i)
		}
		var existing models.Product
		if err := db.Where("sku = ?", sku).First(&existing).Error; err == nil {
			skipped++
			if (i-1)%5000 == 0 && i > 1 {
				fmt.Printf("... %d rows processed\n", i-1)
			}
			continue
		}
		cat, err := catRepo.FindOrCreateByPath(at(row, 0), at(row, 1))
		if err != nil {
			continue
		}
		catID := cat.ID
		stock, _ := strconv.Atoi(strings.TrimSpace(at(row, 7)))
		price1, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, 8)), ",", ".", 1), 64)
		price5, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, 9)), ",", ".", 1), 64)
		price10, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(at(row, 10)), ",", ".", 1), 64)
		imgRaw := strings.TrimSpace(at(row, 14))
		imgPath := ""
		if idx := strings.Index(imgRaw, ","); idx > 0 {
			imgPath = strings.TrimSpace(imgRaw[:idx])
		} else if imgRaw != "" {
			imgPath = imgRaw
		}
		chars := strings.TrimSpace(at(row, 13))
		var attrs []models.ProductAttribute
		if chars != "" {
			attrs = []models.ProductAttribute{{Key: "characteristics", Value: chars}}
		}
		p := models.Product{
			Name:        name,
			SKU:         sku,
			CategoryID:  &catID,
			Description: strings.TrimSpace(at(row, 12)),
			ImagePath:   imgPath,
			Size:        strings.TrimSpace(at(row, 3)),
			Mark:        strings.TrimSpace(at(row, 4)),
			Length:      strings.TrimSpace(at(row, 5)),
			City:        strings.TrimSpace(at(row, 6)),
			Stock:       stock,
			Price1t:     price1,
			Price5t:     price5,
			Price10t:    price10,
			SourceURL:   strings.TrimSpace(at(row, 15)),
			Attributes:  attrs,
		}
		if err := productRepo.Create(&p); err == nil {
			created++
		}
		if (i)%5000 == 0 {
			fmt.Printf("Processed %d rows, created %d, skipped %d\n", i, created, skipped)
		}
	}
	fmt.Printf("Done. Created: %d, Skipped: %d\n", created, skipped)
}

func at(row []string, i int) string {
	if i < len(row) {
		return row[i]
	}
	return ""
}
