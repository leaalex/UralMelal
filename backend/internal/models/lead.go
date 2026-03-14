package models

import (
	"time"

	"gorm.io/gorm"
)

type LeadStatus string

const (
	LeadStatusNew        LeadStatus = "new"
	LeadStatusProcessed  LeadStatus = "processed"
)

type Lead struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:255;not null" json:"name"`
	Phone     string         `gorm:"size:50;not null" json:"phone"`
	Subject   string         `gorm:"size:500" json:"subject"`
	Status    LeadStatus     `gorm:"size:20;default:new" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
