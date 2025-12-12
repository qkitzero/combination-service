package category

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/language"
)

type Category interface {
	ID() CategoryID
	Name() Name
	Language() language.Language
	CreatedAt() time.Time
}

type category struct {
	id        CategoryID
	name      Name
	language  language.Language
	createdAt time.Time
}

func (c category) ID() CategoryID {
	return c.id
}

func (c category) Name() Name {
	return c.name
}

func (c category) Language() language.Language {
	return c.language
}

func (c category) CreatedAt() time.Time {
	return c.createdAt
}

func NewCategory(
	id CategoryID,
	name Name,
	language language.Language,
	createdAt time.Time,
) Category {
	return &category{
		id:        id,
		name:      name,
		language:  language,
		createdAt: createdAt,
	}
}
