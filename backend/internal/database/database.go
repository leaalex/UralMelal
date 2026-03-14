package database

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"uralmelal/backend/internal/config"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch cfg.DBType {
	case "sqlite":
		if err := ensureDir(filepath.Dir(cfg.DBPath)); err != nil {
			return nil, fmt.Errorf("ensure db dir: %w", err)
		}
		dialector = sqlite.Open(cfg.DBPath)
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
		)
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported DB_TYPE: %s", cfg.DBType)
	}

	logLevel := logger.Silent
	if cfg.AppEnv == "dev" {
		logLevel = logger.Info
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	slog.Info("Database connected", "type", cfg.DBType)
	return db, nil
}

func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}
