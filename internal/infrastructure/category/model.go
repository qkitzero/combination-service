package category

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/language"
)

type CategoryModel struct {
	ID           category.CategoryID
	Name         category.Name
	LanguageCode language.Language
	CreatedAt    time.Time
}

func (CategoryModel) TableName() string {
	return "categories"
}
