package rule

import (
	"testing"
)

func TestNewRule(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		success  bool
		count    int
		strategy Strategy
	}{
		{"success new rule", true, 3, NewStrategy(StrategyTypeRandom)},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rule := NewRule(tt.count, tt.strategy)
			if tt.success && rule.Count() != tt.count {
				t.Errorf("Count() = %v, want %v", rule.Count(), tt.count)
			}
			if tt.success && rule.Strategy() != tt.strategy {
				t.Errorf("Strategy() = %v, want %v", rule.Strategy(), tt.strategy)
			}
		})
	}
}
