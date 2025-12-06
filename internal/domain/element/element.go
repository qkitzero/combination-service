package element

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
)

type Element interface {
	ID() ElementID
	Name() Name
	Categories() []category.Category
	CreatedAt() time.Time
}

type element struct {
	id         ElementID
	name       Name
	categories []category.Category
	createdAt  time.Time
}

func (e element) ID() ElementID {
	return e.id
}

func (e element) Name() Name {
	return e.name
}

func (e element) Categories() []category.Category {
	return e.categories
}

func (e element) CreatedAt() time.Time {
	return e.createdAt
}

func NewElement(
	id ElementID,
	name Name,
	categories []category.Category,
	createdAt time.Time,
) Element {
	return &element{
		id:         id,
		name:       name,
		categories: categories,
		createdAt:  createdAt,
	}
}
