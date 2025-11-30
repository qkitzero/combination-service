package element

import "testing"

func TestNewName(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		success bool
		Name    string
	}{
		{"success new name", true, "test user"},
		{"failure empty name", false, ""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			Name, err := NewName(tt.Name)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if tt.success && Name.String() != tt.Name {
				t.Errorf("String() = %v, want %v", Name.String(), tt.Name)
			}
		})
	}
}
