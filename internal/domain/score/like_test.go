package score

import "testing"

func TestNewLike(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		like    int
	}{
		{"success new like", true, 1},
		{"success new like with zero", true, 0},
		{"failure negative like", false, -1},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			like, err := NewLike(tt.like)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if tt.success && like.Int() != tt.like {
				t.Errorf("Int() = %v, want %v", like.Int(), tt.like)
			}
		})
	}
}

func TestIncrement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		success      bool
		like         int
		expectedLike int
	}{
		{"increment from zero", true, 0, 1},
		{"increment from positive", true, 5, 6},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			like, err := NewLike(tt.like)
			if err != nil {
				t.Errorf("failed to new like")
			}

			incrementedLike := like.Increment()

			if incrementedLike.Int() != tt.expectedLike {
				t.Errorf("Increment() = %v, want %v", incrementedLike.Int(), tt.expectedLike)
			}
		})
	}
}
