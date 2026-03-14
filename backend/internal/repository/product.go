package repository

import (
	"strings"

	"uralmelal/backend/internal/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db     *gorm.DB
	dbType string
}

func NewProductRepository(db *gorm.DB, dbType string) *ProductRepository {
	return &ProductRepository{db: db, dbType: dbType}
}

type ProductListParams struct {
	Q          string
	CategoryID *uint
	Page       int
	PerPage    int
	Sort       string
}

type ProductListResult struct {
	Products []models.Product
	Total    int64
}

func (r *ProductRepository) List(params ProductListParams) (ProductListResult, error) {
	q := r.db.Model(&models.Product{})
	if params.CategoryID != nil {
		ids := r.categoryIDsWithDescendants(*params.CategoryID)
		q = q.Where("category_id IN ?", ids)
	}
	if params.Q != "" {
		searchTerm := "%" + strings.TrimSpace(params.Q) + "%"
		// ILIKE для postgres (без учёта регистра), LIKE для sqlite (уже case-insensitive)
		op := "LIKE"
		if r.dbType == "postgres" {
			op = "ILIKE"
		}
		q = q.Where(
			"name "+op+" ? OR sku "+op+" ? OR description "+op+" ? OR size "+op+" ? OR mark "+op+" ? OR city "+op+" ?",
			searchTerm, searchTerm, searchTerm, searchTerm, searchTerm, searchTerm,
		)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return ProductListResult{}, err
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
	order := "created_at DESC"
	switch params.Sort {
	case "name_asc":
		order = "name ASC"
	case "name_desc":
		order = "name DESC"
	case "price_asc":
		order = "price_1t ASC"
	case "price_desc":
		order = "price_1t DESC"
	}
	var products []models.Product
	err := q.Preload("Category").Preload("Attributes").
		Limit(perPage).Offset(offset).
		Order(order).
		Find(&products).Error
	if err != nil {
		return ProductListResult{}, err
	}
	return ProductListResult{Products: products, Total: total}, nil
}

func (r *ProductRepository) categoryIDsWithDescendants(catID uint) []uint {
	ids := []uint{catID}
	var childIDs []uint
	r.db.Model(&models.Category{}).Where("parent_id = ?", catID).Pluck("id", &childIDs)
	for _, cid := range childIDs {
		ids = append(ids, r.categoryIDsWithDescendants(cid)...)
	}
	return ids
}

func (r *ProductRepository) GetByID(id uint) (*models.Product, error) {
	var p models.Product
	err := r.db.Preload("Category").Preload("Attributes").First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *ProductRepository) Create(p *models.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		attrs := p.Attributes
		p.Attributes = nil
		if err := tx.Create(p).Error; err != nil {
			return err
		}
		p.Attributes = attrs
		for i := range p.Attributes {
			p.Attributes[i].ProductID = p.ID
			p.Attributes[i].ID = 0
			if err := tx.Omit("ID").Create(&p.Attributes[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ProductRepository) Update(p *models.Product) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(p).Error; err != nil {
			return err
		}
		tx.Where("product_id = ?", p.ID).Delete(&models.ProductAttribute{})
		for i := range p.Attributes {
			p.Attributes[i].ProductID = p.ID
			p.Attributes[i].ID = 0
			if err := tx.Omit("ID").Create(&p.Attributes[i]).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *ProductRepository) Delete(id uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("product_id = ?", id).Delete(&models.ProductAttribute{}).Error; err != nil {
			return err
		}
		return tx.Delete(&models.Product{}, id).Error
	})
}

