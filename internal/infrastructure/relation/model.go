package relation

import (
	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
)

type ElementCategoryModel struct {
	ElementID  element.ElementID
	CategoryID category.CategoryID
}

func (ElementCategoryModel) TableName() string {
	return "element_category"
}
