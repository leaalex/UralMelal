package models

import (
	"time"

	"gorm.io/gorm"
)

type ContentPage string

const (
	PageHome     ContentPage = "home"
	PageContacts ContentPage = "contacts"
)

type ContentBlock struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Page      ContentPage    `gorm:"size:50;not null;index" json:"page"`
	Slug      string         `gorm:"size:255;not null;index" json:"slug"`
	Title     string         `gorm:"size:500" json:"title"`
	Body      string         `gorm:"type:text" json:"body"`
	SortOrder int            `gorm:"default:0" json:"sort_order"`
	Hidden    bool           `gorm:"default:false" json:"hidden"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
