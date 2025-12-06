package category

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
)

type CategoryModel struct {
	ID        category.CategoryID
	Name      category.Name
	CreatedAt time.Time
}

func (CategoryModel) TableName() string {
	return "categories"
}
