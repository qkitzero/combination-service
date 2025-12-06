package combination

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
)

func TestCreateElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		success     bool
		elementName string
		createErr   error
	}{
		{"success create element", true, "test element", nil},
		{"failure empty name", false, "", nil},
		{"failure create error", false, "test element", errors.New("create error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockElementRepository.EXPECT().Create(gomock.Any()).Return(tt.createErr).AnyTimes()
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.CreateElement(tt.elementName)
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
		name         string
		success      bool
		categoryName string
		createErr    error
	}{
		{"success create category", true, "test category", nil},
		{"failure empty name", false, "", nil},
		{"failure create error", false, "test category", errors.New("create error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)
			mockCategoryRepository.EXPECT().Create(gomock.Any()).Return(tt.createErr).AnyTimes()

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.CreateCategory(tt.categoryName)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
