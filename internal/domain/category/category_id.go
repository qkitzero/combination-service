package category

import (
	"fmt"

	"github.com/google/uuid"
)

type CategoryID struct {
	uuid.UUID
}

func NewCategoryID() CategoryID {
	id := uuid.New()
	return CategoryID{id}
}

func NewCategoryIDFromString(s string) (CategoryID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return CategoryID{}, fmt.Errorf("invalid UUID format: %w", err)
	}
	return CategoryID{id}, nil
}
