package element

import (
	"time"
)

type Element interface {
	ID() ElementID
	Name() Name
	CreatedAt() time.Time
}

type element struct {
	id        ElementID
	name      Name
	createdAt time.Time
}

func (e element) ID() ElementID {
	return e.id
}

func (e element) Name() Name {
	return e.name
}

func (e element) CreatedAt() time.Time {
	return e.createdAt
}

func NewElement(
	id ElementID,
	name Name,
	createdAt time.Time,
) Element {
	return &element{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}
}
