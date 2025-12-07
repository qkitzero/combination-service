package combination

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/qkitzero/combination-service/internal/domain/category"
	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
)

func TestCreateElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		success     bool
		elementName string
		categoryIDs []string
		createErr   error
		findByIDErr error
	}{
		{"success create element", true, "test element", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, nil},
		{"success create element no category", true, "test element", []string{}, nil, nil},
		{"failure empty name", false, "", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, nil},
		{"failure create error", false, "test element", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, errors.New("create error"), nil},
		{"failure invalid category id", false, "test element", []string{"0123456789"}, nil, nil},
		{"failure find by id error", false, "test element", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, errors.New("find by id error")},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockElementRepository.EXPECT().Create(gomock.Any()).Return(tt.createErr).AnyTimes()
			mockCategories := make([]category.Category, len(tt.categoryIDs))
			for i := range mockCategories {
				mockCategory := mockscategory.NewMockCategory(ctrl)
				mockCategories[i] = mockCategory
			}
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)
			mockCategoryRepository.EXPECT().FindAllByIDs(gomock.Any()).Return(mockCategories, tt.findByIDErr).AnyTimes()

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.CreateElement(tt.elementName, tt.categoryIDs)
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
