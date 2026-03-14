package repository

import (
	"time"

	"uralmelal/backend/internal/models"

	"gorm.io/gorm"
)

type LeadListParams struct {
	Status    *models.LeadStatus
	DateFrom  *time.Time
	DateTo    *time.Time
	Page      int
	PerPage   int
}

type LeadRepository struct {
	db *gorm.DB
}

func NewLeadRepository(db *gorm.DB) *LeadRepository {
	return &LeadRepository{db: db}
}

func (r *LeadRepository) Create(l *models.Lead) error {
	return r.db.Create(l).Error
}

func (r *LeadRepository) List(params LeadListParams) ([]models.Lead, int64, error) {
	q := r.db.Model(&models.Lead{})
	if params.Status != nil {
		q = q.Where("status = ?", *params.Status)
	}
	if params.DateFrom != nil {
		q = q.Where("created_at >= ?", *params.DateFrom)
	}
	if params.DateTo != nil {
		q = q.Where("created_at <= ?", *params.DateTo)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	perPage := params.PerPage
	if perPage <= 0 {
		perPage = 20
	}
	if perPage > 100 {
		perPage = 100
	}
	page := params.Page
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * perPage
	var leads []models.Lead
	err := q.Order("created_at DESC").Limit(perPage).Offset(offset).Find(&leads).Error
	return leads, total, err
}

func (r *LeadRepository) GetByID(id uint) (*models.Lead, error) {
	var l models.Lead
	err := r.db.First(&l, id).Error
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *LeadRepository) UpdateStatus(id uint, status models.LeadStatus) error {
	return r.db.Model(&models.Lead{}).Where("id = ?", id).Update("status", status).Error
}
