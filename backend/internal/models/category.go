package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Slug      string         `gorm:"uniqueIndex;size:255;not null" json:"slug"`
	ParentID  *uint          `json:"parent_id"`
	Parent    *Category      `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children  []Category    `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	SortOrder int            `gorm:"default:0" json:"sort_order"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
