package element

import (
	"fmt"

	"github.com/google/uuid"
)

type ElementID struct {
	uuid.UUID
}

func NewElementID() ElementID {
	id := uuid.New()
	return ElementID{id}
}

func NewElementIDFromString(s string) (ElementID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ElementID{}, fmt.Errorf("invalid UUID format: %w", err)
	}
	return ElementID{id}, nil
}
