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
		{"success new strategy with unknown type", true, "unknown"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			strategy := NewStrategy(tt.strategyType)
			if tt.success && strategy == nil {
				t.Errorf("NewStrategy() = nil, want non-nil")
			}
		})
	}
}
