package language

import (
	"testing"
)

func TestNewLanguage(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		success      bool
		languageCode string
	}{
		{"success new language en", true, "en"},
		{"success new language ja", true, "ja"},
		{"failure empty language code", false, ""},
		{"failure invalid language code", false, "english"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			language, err := NewLanguage(tt.languageCode)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if tt.success && language.String() != tt.languageCode {
				t.Errorf("String() = %v, want %v", language.String(), tt.languageCode)
			}
		})
	}
}
