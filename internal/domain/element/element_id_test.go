package element

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewElementID(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
	}{
		{"success new element id", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			id := NewElementID()
			if tt.success && id.UUID == uuid.Nil {
				t.Errorf("expected valid element id, but got a nil UUID")
			}
		})
	}
}

func TestNewElementIDFromString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		id      string
	}{
		{"success new element id from string", true, "fe8c2263-bbac-4bb9-a41d-b04f5afc4425"},
		{"failure empty element id", false, ""},
		{"failure invalid element id", false, "0123456789"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewElementIDFromString(tt.id)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
