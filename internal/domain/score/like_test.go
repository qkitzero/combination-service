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
