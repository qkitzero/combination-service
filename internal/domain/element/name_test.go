package element

import "testing"

func TestNewName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		success     bool
		elementName string
	}{
		{"success new name", true, "test name"},
		{"failure empty name", false, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			elementName, err := NewName(tt.elementName)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if tt.success && elementName.String() != tt.elementName {
				t.Errorf("String() = %v, want %v", elementName.String(), tt.elementName)
			}
		})
	}
}
