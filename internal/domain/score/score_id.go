package score

import (
	"fmt"

	"github.com/google/uuid"
)

type ScoreID struct {
	uuid.UUID
}

func NewScoreID() ScoreID {
	id := uuid.New()
	return ScoreID{id}
}

func NewScoreIDFromString(s string) (ScoreID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return ScoreID{}, fmt.Errorf("invalid UUID format: %w", err)
	}
	return ScoreID{id}, nil
}
