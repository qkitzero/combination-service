package element

import (
	"reflect"
	"testing"
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/language"
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
	categoryID, err := category.NewCategoryIDFromString("bcf148ed-5a7f-45c2-95d8-c190179b548e")
	if err != nil {
		t.Errorf("failed to new category id: %v", err)
	}
	elementLanguage, err := language.NewLanguage("en")
	if err != nil {
		t.Errorf("failed to new language: %v", err)
	}
	categoryName, err := category.NewName("test category")
	if err != nil {
		t.Errorf("failed to new category name: %v", err)
	}
	categoryLanguage, err := language.NewLanguage("en")
	if err != nil {
		t.Errorf("failed to new language: %v", err)
	}
	categories := []category.Category{category.NewCategory(categoryID, categoryName, categoryLanguage, time.Now())}
	tests := []struct {
		name            string
		success         bool
		id              ElementID
		elementName     Name
		elementLanguage language.Language
		categories      []category.Category
		createdAt       time.Time
	}{
		{"success new element", true, id, elementName, elementLanguage, categories, time.Now()},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			element := NewElement(tt.id, tt.elementName, tt.elementLanguage, tt.categories, tt.createdAt)
			if tt.success && element.ID() != tt.id {
				t.Errorf("ID() = %v, want %v", element.ID(), tt.id)
			}
			if tt.success && element.Name() != tt.elementName {
				t.Errorf("Name() = %v, want %v", element.Name(), tt.elementName)
			}
			if tt.success && element.Language() != tt.elementLanguage {
				t.Errorf("Language() = %v, want %v", element.Language(), tt.elementLanguage)
			}
			if tt.success && !reflect.DeepEqual(element.Categories(), tt.categories) {
				t.Errorf("Categories() = %v, want %v", element.Categories(), tt.categories)
			}
			if tt.success && !element.CreatedAt().Equal(tt.createdAt) {
				t.Errorf("CreatedAt() = %v, want %v", element.CreatedAt(), tt.createdAt)
			}
		})
	}
}
