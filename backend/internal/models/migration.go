package models

import (
	"time"

	"gorm.io/gorm"
)

type AppliedMigration struct {
	Name      string         `gorm:"primaryKey;size:100" json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
