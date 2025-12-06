package category

import (
	"gorm.io/gorm"

	"github.com/qkitzero/combination-service/internal/domain/category"
)

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(c category.Category) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		categoryModel := CategoryModel{
			ID:        c.ID(),
			Name:      c.Name(),
			CreatedAt: c.CreatedAt(),
		}

		if err := tx.Create(&categoryModel).Error; err != nil {
			return err
		}

		return nil
	})
}
