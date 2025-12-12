package element

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
	"github.com/qkitzero/combination-service/internal/domain/language"
)

type ElementModel struct {
	ID           element.ElementID
	Name         element.Name
	LanguageCode language.Language
	CreatedAt    time.Time
}

func (ElementModel) TableName() string {
	return "elements"
}
