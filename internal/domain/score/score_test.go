package score

import (
	"testing"
	"time"
)

func TestNewScore(t *testing.T) {
	t.Parallel()
	id, err := NewScoreIDFromString("9a356e0c-e0a3-413e-8a4e-c4b6d435511d")
	if err != nil {
		t.Errorf("failed to new score id: %v", err)
	}
	like, err := NewLike(1)
	if err != nil {
		t.Errorf("failed to new like: %v", err)
	}
	tests := []struct {
		name      string
		success   bool
		id        ScoreID
		like      Like
		createdAt time.Time
		updatedAt time.Time
	}{
		{"success new score", true, id, like, time.Now(), time.Now()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			score := NewScore(tt.id, tt.like, tt.createdAt, tt.updatedAt)
			if tt.success && score.ID() != tt.id {
				t.Errorf("ID() = %v, want %v", score.ID(), tt.id)
			}
			if tt.success && score.Like() != tt.like {
				t.Errorf("Like() = %v, want %v", score.Like(), tt.like)
			}
			if tt.success && !score.CreatedAt().Equal(tt.createdAt) {
				t.Errorf("CreatedAt() = %v, want %v", score.CreatedAt(), tt.createdAt)
			}
			if tt.success && !score.UpdatedAt().Equal(tt.updatedAt) {
				t.Errorf("UpdatedAt() = %v, want %v", score.UpdatedAt(), tt.updatedAt)
			}
		})
	}
}
