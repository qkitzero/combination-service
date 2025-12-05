package category

import (
	"time"
)

type Category interface {
	ID() CategoryID
	Name() Name
	CreatedAt() time.Time
}

type category struct {
	id        CategoryID
	name      Name
	createdAt time.Time
}

func (c category) ID() CategoryID {
	return c.id
}

func (c category) Name() Name {
	return c.name
}

func (c category) CreatedAt() time.Time {
	return c.createdAt
}

func NewCategory(
	id CategoryID,
	name Name,
	createdAt time.Time,
) Category {
	return &category{
		id:        id,
		name:      name,
		createdAt: createdAt,
	}
}
