package category

import (
	"errors"

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

func (r *categoryRepository) FindByID(id category.CategoryID) (category.Category, error) {
	var categoryModel CategoryModel
	err := r.db.Where("id = ?", id).First(&categoryModel).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, category.ErrCategoryNotFound
	}
	if err != nil {
		return nil, err
	}

	return category.NewCategory(
		categoryModel.ID,
		categoryModel.Name,
		categoryModel.CreatedAt,
	), nil
}

func (r *categoryRepository) FindAllByIDs(ids []category.CategoryID) ([]category.Category, error) {
	if len(ids) == 0 {
		return []category.Category{}, nil
	}

	var categoryModels []CategoryModel
	err := r.db.Where("id IN ?", ids).Find(&categoryModels).Error
	if err != nil {
		return nil, err
	}

	categories := make([]category.Category, len(categoryModels))
	for i, m := range categoryModels {
		categories[i] = category.NewCategory(
			m.ID,
			m.Name,
			m.CreatedAt,
		)
	}

	return categories, nil
}
