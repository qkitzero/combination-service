package element

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type ElementModel struct {
	ID        element.ElementID
	Name      element.Name
	CreatedAt time.Time
}

func (ElementModel) TableName() string {
	return "elements"
}
