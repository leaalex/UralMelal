package repository

import (
	"uralmelal/backend/internal/models"
	"uralmelal/backend/internal/util"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetTree() ([]models.Category, error) {
	var roots []models.Category
	err := r.db.Where("parent_id IS NULL").Order("sort_order ASC, name ASC").Preload("Children", func(db *gorm.DB) *gorm.DB {
		return db.Order("sort_order ASC, name ASC")
	}).Find(&roots).Error
	return roots, err
}

func (r *CategoryRepository) List() ([]models.Category, error) {
	var cats []models.Category
	err := r.db.Order("sort_order ASC, name ASC").Find(&cats).Error
	return cats, err
}

func (r *CategoryRepository) GetByID(id uint) (*models.Category, error) {
	var c models.Category
	err := r.db.First(&c, id).Error
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (r *CategoryRepository) Create(c *models.Category) error {
	return r.db.Create(c).Error
}

func (r *CategoryRepository) Update(c *models.Category) error {
	return r.db.Save(c).Error
}

func (r *CategoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

// FindOrCreateByPath finds or creates category by parent name and child name (2-level tree).
// Returns the leaf category ID (подкатегория). If child is empty, returns parent.
func (r *CategoryRepository) FindOrCreateByPath(parentName, childName string) (*models.Category, error) {
	parentSlug := util.Slugify(parentName)
	if parentSlug == "" {
		parentSlug = "uncategorized"
	}

	var parent models.Category
	err := r.db.Where("slug = ? AND parent_id IS NULL", parentSlug).First(&parent).Error
	if err != nil {
		parent = models.Category{Name: parentName, Slug: parentSlug}
		if err := r.db.Create(&parent).Error; err != nil {
			return nil, err
		}
	}

	if childName == "" {
		return &parent, nil
	}

	childSlug := parentSlug + "-" + util.Slugify(childName)
	if util.Slugify(childName) == "" {
		childSlug = parentSlug + "-child"
	}
	pid := parent.ID

	var child models.Category
	err = r.db.Where("slug = ? AND parent_id = ?", childSlug, pid).First(&child).Error
	if err != nil {
		child = models.Category{Name: childName, Slug: childSlug, ParentID: &pid}
		if err := r.db.Create(&child).Error; err != nil {
			return nil, err
		}
	}
	return &child, nil
}
