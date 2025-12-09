package combination

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/google/uuid"
	combinationv1 "github.com/qkitzero/combination-service/gen/go/combination/v1"
	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
	mocksappcombination "github.com/qkitzero/combination-service/mocks/application/combination"
	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
)

func TestCreateElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		success          bool
		ctx              context.Context
		elementName      string
		categoryIDs      []string
		createElementErr error
	}{
		{"success create element", true, context.Background(), "test element", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil},
		{"failure create element error", false, context.Background(), "test element", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, fmt.Errorf("create element error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElement := mockselement.NewMockElement(ctrl)
			mockElement.EXPECT().ID().Return(element.ElementID{UUID: uuid.New()}).AnyTimes()
			mockCombinationUsecase := mocksappcombination.NewMockCombinationUsecase(ctrl)
			mockCombinationUsecase.EXPECT().CreateElement(tt.elementName, tt.categoryIDs).Return(mockElement, tt.createElementErr).AnyTimes()

			combinationHandler := NewCombinationHandler(mockCombinationUsecase)

			req := &combinationv1.CreateElementRequest{
				Name:        tt.elementName,
				CategoryIds: tt.categoryIDs,
			}

			_, err := combinationHandler.CreateElement(tt.ctx, req)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}

func TestListElements(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name            string
		success         bool
		ctx             context.Context
		listElementsErr error
	}{
		{"success list elements", true, context.Background(), nil},
		{"failure list elements error", false, context.Background(), fmt.Errorf("list elements error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCategory := mockscategory.NewMockCategory(ctrl)
			mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
			mockCategory.EXPECT().Name().Return(category.Name("test category")).AnyTimes()
			mockElement := mockselement.NewMockElement(ctrl)
			mockElement.EXPECT().ID().Return(element.ElementID{UUID: uuid.New()}).AnyTimes()
			mockElement.EXPECT().Name().Return(element.Name("test element")).AnyTimes()
			mockElement.EXPECT().Categories().Return([]category.Category{mockCategory}).AnyTimes()
			mockCombinationUsecase := mocksappcombination.NewMockCombinationUsecase(ctrl)
			mockCombinationUsecase.EXPECT().ListElements().Return([]element.Element{mockElement}, tt.listElementsErr).AnyTimes()

			combinationHandler := NewCombinationHandler(mockCombinationUsecase)

			req := &combinationv1.ListElementsRequest{}

			_, err := combinationHandler.ListElements(tt.ctx, req)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}

func TestCreateCategory(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		success           bool
		ctx               context.Context
		categoryName      string
		createCategoryErr error
	}{
		{"success create category", true, context.Background(), "test category", nil},
		{"failure create category error", false, context.Background(), "test category", fmt.Errorf("create category error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCategory := mockscategory.NewMockCategory(ctrl)
			mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
			mockCombinationUsecase := mocksappcombination.NewMockCombinationUsecase(ctrl)
			mockCombinationUsecase.EXPECT().CreateCategory(tt.categoryName).Return(mockCategory, tt.createCategoryErr).AnyTimes()

			combinationHandler := NewCombinationHandler(mockCombinationUsecase)

			req := &combinationv1.CreateCategoryRequest{
				Name: tt.categoryName,
			}

			_, err := combinationHandler.CreateCategory(tt.ctx, req)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}

func TestListCategories(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name              string
		success           bool
		ctx               context.Context
		listCategoriesErr error
	}{
		{"success list categories", true, context.Background(), nil},
		{"failure list categories error", false, context.Background(), fmt.Errorf("list categories error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockCategory := mockscategory.NewMockCategory(ctrl)
			mockCategory.EXPECT().ID().Return(category.CategoryID{UUID: uuid.New()}).AnyTimes()
			mockCategory.EXPECT().Name().Return(category.Name("test category")).AnyTimes()
			mockCombinationUsecase := mocksappcombination.NewMockCombinationUsecase(ctrl)
			mockCombinationUsecase.EXPECT().ListCategories().Return([]category.Category{mockCategory}, tt.listCategoriesErr).AnyTimes()

			combinationHandler := NewCombinationHandler(mockCombinationUsecase)

			req := &combinationv1.ListCategoriesRequest{}

			_, err := combinationHandler.ListCategories(tt.ctx, req)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
