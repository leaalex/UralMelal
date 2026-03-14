package database

import (
	"log/slog"
	"os"

	"uralmelal/backend/internal/config"
	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/repository"
	"uralmelal/backend/internal/services"

	"gorm.io/gorm"
)

func SeedAdmin(db *gorm.DB, cfg *config.Config) error {
	userRepo := repository.NewUserRepository(db)
	adminUser := getEnvSeed("ADMIN_USERNAME", "admin")
	adminPass := getEnvSeed("ADMIN_PASSWORD", "admin123")

	exists, err := userRepo.ExistsByUsername(adminUser)
	if err != nil {
		return err
	}
	if exists {
		slog.Debug("Admin user already exists, skipping seed")
		return nil
	}

	hash, err := services.HashPassword(adminPass)
	if err != nil {
		return err
	}
	user := &models.User{
		Username:     adminUser,
		PasswordHash: hash,
		Role:         models.RoleAdmin,
	}
	if err := userRepo.Create(user); err != nil {
		return err
	}
	slog.Info("Created admin user", "username", adminUser)
	return nil
}

func SeedContentBlocks(db *gorm.DB) error {
	var count int64
	db.Model(&models.ContentBlock{}).Count(&count)
	if count == 0 {
		blocks := []models.ContentBlock{
			{Page: models.PageHome, Slug: "banner", Title: "Баннер", Body: "", SortOrder: 0, Hidden: true},
			{Page: models.PageHome, Slug: "about", Title: "О компании", Body: "", SortOrder: 1, Hidden: true},
			{Page: models.PageContacts, Slug: "info", Title: "Контактная информация", Body: "", SortOrder: 0},
		}
		for _, b := range blocks {
			db.Create(&b)
		}
		slog.Info("Created default content blocks")
		return nil
	}
	// Миграция: скрыть Баннер и О компании по умолчанию для существующих БД
	var applied int64
	db.Model(&models.AppliedMigration{}).Where("name = ?", "content_blocks_hidden_default").Count(&applied)
	if applied == 0 {
		db.Model(&models.ContentBlock{}).Where("slug IN ?", []string{"banner", "about"}).Update("hidden", true)
		db.Create(&models.AppliedMigration{Name: "content_blocks_hidden_default"})
		slog.Info("Applied migration: content_blocks_hidden_default")
	}
	return nil
}

func getEnvSeed(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
