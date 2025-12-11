package rule_test

import (
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/qkitzero/combination-service/internal/domain/element"
	"github.com/qkitzero/combination-service/internal/domain/rule"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
	mocksrule "github.com/qkitzero/combination-service/mocks/domain/rule"
)

func TestNewRule(t *testing.T) {
	t.Parallel()
	strategy, err := rule.NewStrategy(rule.StrategyTypeRandom)
	if err != nil {
		t.Errorf("failed to new strategy: %v", err)
	}
	tests := []struct {
		name     string
		success  bool
		count    int
		strategy rule.Strategy
	}{
		{"success new rule", true, 3, strategy},
		{"failure negative count", false, -1, strategy},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			rule, err := rule.NewRule(tt.count, tt.strategy)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}

			if tt.success && rule.Count() != tt.count {
				t.Errorf("Count() = %v, want %v", rule.Count(), tt.count)
			}
			if tt.success && rule.Strategy() != tt.strategy {
				t.Errorf("Strategy() = %v, want %v", rule.Strategy(), tt.strategy)
			}
		})
	}
}

func TestApply(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		success     bool
		count       int
		numElements int
	}{
		{"success apply", true, 3, 5},
		{"success zero count", true, 0, 5},
		{"failure not enough elements", false, 3, 2},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockStrategy := mocksrule.NewMockStrategy(ctrl)
			mockStrategy.EXPECT().Select(tt.count, gomock.Any()).Return([]element.Element{}, nil).AnyTimes()

			rule, err := rule.NewRule(tt.count, mockStrategy)
			if err != nil {
				t.Errorf("failed to new rule: %v", err)
			}

			mockElement := mockselement.NewMockElement(ctrl)

			mockElements := make([]element.Element, tt.numElements)
			for i := range tt.numElements {
				mockElements[i] = mockElement
			}

			_, err = rule.Apply(mockElements)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
