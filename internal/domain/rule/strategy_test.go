package rule

import (
	"testing"
)

func TestNewStrategy(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		success      bool
		strategyType StrategyType
	}{
		{"success new strategy", true, "random"},
		{"failure unknown type", false, "unknown"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := NewStrategy(tt.strategyType)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
