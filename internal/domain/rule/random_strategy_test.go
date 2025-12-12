package rule

import (
	"fmt"
	"testing"
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
	"github.com/qkitzero/combination-service/internal/domain/language"
)

func TestRandomStrategySelect(t *testing.T) {
	t.Parallel()
	numElements := 5
	elements := make([]element.Element, numElements)
	for i := range numElements {
		elementName, err := element.NewName(fmt.Sprintf("element%d", i+1))
		if err != nil {
			t.Fatalf("failed to new element name: %v", err)
		}
		elementLanguage, err := language.NewLanguage("en")
		if err != nil {
			t.Fatalf("failed to new element language: %v", err)
		}
		elements[i] = element.NewElement(element.NewElementID(), elementName, elementLanguage, []category.Category{}, time.Now())
	}
	tests := []struct {
		name     string
		success  bool
		count    int
		elements []element.Element
	}{
		{"success random strategy select", true, 3, elements},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			randomStrategy := newRandomStrategy()

			_, err := randomStrategy.Select(tt.count, tt.elements)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
