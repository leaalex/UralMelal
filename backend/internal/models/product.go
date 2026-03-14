package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	Name        string            `gorm:"size:500;not null;index" json:"name"`
	SKU         string            `gorm:"size:100;index" json:"sku"`
	CategoryID  *uint             `json:"category_id"`
	Category    *Category         `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	Description string            `gorm:"type:text" json:"description"`
	ImagePath   string            `gorm:"size:500" json:"image_path"`
	// MC import fields
	Size      string  `gorm:"size:100" json:"size"`
	Mark      string  `gorm:"size:100" json:"mark"`
	Length    string  `gorm:"size:100" json:"length"`
	City      string  `gorm:"size:100" json:"city"`
	Stock     int     `gorm:"default:0" json:"stock"`
	Price1t   float64 `gorm:"type:decimal(12,2)" json:"price_1t"`
	Price5t   float64 `gorm:"type:decimal(12,2)" json:"price_5t"`
	Price10t  float64 `gorm:"type:decimal(12,2)" json:"price_10t"`
	SourceURL string  `gorm:"size:500" json:"source_url"`
	Attributes  []ProductAttribute `gorm:"foreignKey:ProductID" json:"attributes,omitempty"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"-"`
}

type ProductAttribute struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	ProductID uint           `gorm:"not null;index" json:"product_id"`
	Key       string         `gorm:"size:255;not null" json:"key"`
	Value     string         `gorm:"size:500" json:"value"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
