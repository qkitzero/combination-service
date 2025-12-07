package element

import (
	"gorm.io/gorm"

	"github.com/qkitzero/combination-service/internal/domain/element"
	"github.com/qkitzero/combination-service/internal/infrastructure/relation"
)

type elementRepository struct {
	db *gorm.DB
}

func NewElementRepository(db *gorm.DB) element.ElementRepository {
	return &elementRepository{db: db}
}

func (r *elementRepository) Create(e element.Element) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		elementModel := ElementModel{
			ID:        e.ID(),
			Name:      e.Name(),
			CreatedAt: e.CreatedAt(),
		}

		if err := tx.Create(&elementModel).Error; err != nil {
			return err
		}

		if len(e.Categories()) == 0 {
			return nil
		}

		var elementCategoryModels []relation.ElementCategoryModel
		for _, c := range e.Categories() {
			elementCategoryModels = append(elementCategoryModels, relation.ElementCategoryModel{
				ElementID:  e.ID(),
				CategoryID: c.ID(),
			})
		}

		if err := tx.Create(&elementCategoryModels).Error; err != nil {
			return err
		}

		return nil
	})
}
