package category

import (
	"testing"
	"time"
)

func TestNewCategory(t *testing.T) {
	t.Parallel()
	id, err := NewCategoryIDFromString("fe8c2263-bbac-4bb9-a41d-b04f5afc4425")
	if err != nil {
		t.Errorf("failed to new category id: %v", err)
	}
	categoryName, err := NewName("test category")
	if err != nil {
		t.Errorf("failed to new name: %v", err)
	}
	tests := []struct {
		name         string
		success      bool
		id           CategoryID
		categoryName Name
		createdAt    time.Time
	}{
		{"success new category", true, id, categoryName, time.Now()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			category := NewCategory(tt.id, tt.categoryName, tt.createdAt)
			if tt.success && category.ID() != tt.id {
				t.Errorf("ID() = %v, want %v", category.ID(), tt.id)
			}
			if tt.success && category.Name() != tt.categoryName {
				t.Errorf("Name() = %v, want %v", category.Name(), tt.categoryName)
			}
			if tt.success && !category.CreatedAt().Equal(tt.createdAt) {
				t.Errorf("CreatedAt() = %v, want %v", category.CreatedAt(), tt.createdAt)
			}

		})
	}
}
