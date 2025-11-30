package element

import (
	"testing"
	"time"
)

func TestNewElement(t *testing.T) {
	t.Parallel()
	id, err := NewElementIDFromString("fe8c2263-bbac-4bb9-a41d-b04f5afc4425")
	if err != nil {
		t.Errorf("failed to new element id: %v", err)
	}
	elementName, err := NewName("test element")
	if err != nil {
		t.Errorf("failed to new name: %v", err)
	}
	tests := []struct {
		name        string
		success     bool
		id          ElementID
		elementName Name
		createdAt   time.Time
	}{
		{"success new element", true, id, elementName, time.Now()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			element := NewElement(tt.id, tt.elementName, tt.createdAt)
			if tt.success && element.ID() != tt.id {
				t.Errorf("ID() = %v, want %v", element.ID(), tt.id)
			}
			if tt.success && element.Name() != tt.elementName {
				t.Errorf("Name() = %v, want %v", element.Name(), tt.elementName)
			}
			if tt.success && !element.CreatedAt().Equal(tt.createdAt) {
				t.Errorf("CreatedAt() = %v, want %v", element.CreatedAt(), tt.createdAt)
			}

		})
	}
}
