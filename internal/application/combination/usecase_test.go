package combination

import (
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
	mockscategory "github.com/qkitzero/combination-service/mocks/domain/category"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
)

func TestCreateElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name         string
		success      bool
		elementName  string
		languageCode string
		categoryIDs  []string
		createErr    error
		findByIDErr  error
	}{
		{"success create element", true, "test element", "en", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, nil},
		{"success create element no category", true, "test element", "en", []string{}, nil, nil},
		{"failure empty name", false, "", "en", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, nil},
		{"failure empty language code", false, "test element", "", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, nil},
		{"failure create error", false, "test element", "en", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, errors.New("create error"), nil},
		{"failure invalid category id", false, "test element", "en", []string{"0123456789"}, nil, nil},
		{"failure find by id error", false, "test element", "en", []string{"91b349ab-2ffc-45cd-adab-61d248b3f9d9"}, nil, errors.New("find by id error")},
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

			_, err := combinationUsecase.CreateElement(tt.elementName, tt.languageCode, tt.categoryIDs)
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
		name    string
		success bool
	}{
		{"success list elements", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockElementRepository.EXPECT().FindAll().Return([]element.Element{}, nil).AnyTimes()
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.ListElements()
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
		languageCode string
		createErr    error
	}{
		{"success create category", true, "test category", "en", nil},
		{"failure empty name", false, "", "en", nil},
		{"failure empty language code", false, "test category", "", nil},
		{"failure create error", false, "test category", "en", errors.New("create error")},
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

			_, err := combinationUsecase.CreateCategory(tt.categoryName, tt.languageCode)
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
		name    string
		success bool
	}{
		{"success list categories", true},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)
			mockCategoryRepository.EXPECT().FindAll().Return([]category.Category{}, nil).AnyTimes()

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.ListCategories()
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}

func TestGetCombination(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name       string
		success    bool
		count      int
		findAllErr error
	}{
		{"success get combination", true, 3, nil},
		{"failure invalid count", false, -1, nil},
		{"failure find all error", false, 3, errors.New("find all error")},
		{"failure not enough elements", false, 10, nil},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockElement := mockselement.NewMockElement(ctrl)
			mockElementRepository := mockselement.NewMockElementRepository(ctrl)
			mockElementRepository.EXPECT().FindAll().Return([]element.Element{mockElement, mockElement, mockElement, mockElement, mockElement}, tt.findAllErr).AnyTimes()
			mockCategoryRepository := mockscategory.NewMockCategoryRepository(ctrl)

			combinationUsecase := NewCombinationUsecase(mockElementRepository, mockCategoryRepository)

			_, err := combinationUsecase.GetCombination(tt.count)
			if tt.success && err != nil {
				t.Errorf("expected no error, but got %v", err)
			}
			if !tt.success && err == nil {
				t.Errorf("expected error, but got nil")
			}
		})
	}
}
