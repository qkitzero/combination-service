package score

import (
	"time"
)

type Score interface {
	ID() ScoreID
	Like() Like
	CreatedAt() time.Time
	UpdatedAt() time.Time
	IncrementLike()
}

type score struct {
	id        ScoreID
	like      Like
	createdAt time.Time
	updatedAt time.Time
}

func (s score) ID() ScoreID {
	return s.id
}

func (s score) Like() Like {
	return s.like
}

func (s score) CreatedAt() time.Time {
	return s.createdAt
}

func (s score) UpdatedAt() time.Time {
	return s.updatedAt
}

func (s *score) IncrementLike() {
	s.like = s.like.Increment()
	s.updatedAt = time.Now()
}

func NewScore(
	id ScoreID,
	like Like,
	createdAt time.Time,
	updatedAt time.Time,
) Score {
	return &score{
		id:        id,
		like:      like,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}
