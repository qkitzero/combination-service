package element

import (
	"gorm.io/gorm"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
	infracategory "github.com/qkitzero/combination-service/internal/infrastructure/category"
	infrarelation "github.com/qkitzero/combination-service/internal/infrastructure/relation"
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

		elementCategoryModels := make([]infrarelation.ElementCategoryModel, len(e.Categories()))
		for i, c := range e.Categories() {
			elementCategoryModels[i] = infrarelation.ElementCategoryModel{
				ElementID:  e.ID(),
				CategoryID: c.ID(),
			}
		}

		if err := tx.Create(&elementCategoryModels).Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *elementRepository) FindAll() ([]element.Element, error) {
	var elementModels []ElementModel
	if err := r.db.Find(&elementModels).Error; err != nil {
		return nil, err
	}

	elementIDs := make([]element.ElementID, len(elementModels))
	for i, e := range elementModels {
		elementIDs[i] = e.ID
	}

	var elementCategoryModels []infrarelation.ElementCategoryModel
	if err := r.db.Where("element_id IN ?", elementIDs).Find(&elementCategoryModels).Error; err != nil {
		return nil, err
	}

	categoryIDs := make([]category.CategoryID, 0)
	for _, ec := range elementCategoryModels {
		categoryIDs = append(categoryIDs, ec.CategoryID)
	}

	var categoryModels []infracategory.CategoryModel
	if err := r.db.Where("id IN ?", categoryIDs).Find(&categoryModels).Error; err != nil {
		return nil, err
	}

	categoryMap := make(map[category.CategoryID]infracategory.CategoryModel)
	for _, c := range categoryModels {
		categoryMap[c.ID] = c
	}

	elementCategories := make(map[element.ElementID][]category.Category)
	for _, ec := range elementCategoryModels {
		c, ok := categoryMap[ec.CategoryID]
		if !ok {
			continue
		}
		elementCategories[ec.ElementID] = append(
			elementCategories[ec.ElementID],
			category.NewCategory(
				c.ID,
				c.Name,
				c.CreatedAt,
			),
		)
	}

	elements := make([]element.Element, len(elementModels))
	for i, e := range elementModels {
		elements[i] = element.NewElement(
			e.ID,
			e.Name,
			elementCategories[e.ID],
			e.CreatedAt,
		)
	}

	return elements, nil
}
