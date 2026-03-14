package repository

import (
	"uralmelal/backend/internal/models"

	"gorm.io/gorm"
)

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) *ContentRepository {
	return &ContentRepository{db: db}
}

func (r *ContentRepository) GetByPage(page models.ContentPage) ([]models.ContentBlock, error) {
	var blocks []models.ContentBlock
	err := r.db.Where("page = ? AND hidden = ?", page, false).Order("sort_order ASC, id ASC").Find(&blocks).Error
	return blocks, err
}

func (r *ContentRepository) GetByPageAll(page models.ContentPage) ([]models.ContentBlock, error) {
	var blocks []models.ContentBlock
	err := r.db.Where("page = ?", page).Order("sort_order ASC, id ASC").Find(&blocks).Error
	return blocks, err
}

func (r *ContentRepository) GetByID(id uint) (*models.ContentBlock, error) {
	var b models.ContentBlock
	err := r.db.First(&b, id).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *ContentRepository) Update(b *models.ContentBlock) error {
	return r.db.Save(b).Error
}

func (r *ContentRepository) Create(b *models.ContentBlock) error {
	return r.db.Create(b).Error
}
