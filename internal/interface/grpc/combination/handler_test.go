package combination

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/google/uuid"
	combinationv1 "github.com/qkitzero/combination-service/gen/go/combination/v1"
	"github.com/qkitzero/combination-service/internal/domain/element"
	mocksappcombination "github.com/qkitzero/combination-service/mocks/application/combination"
	mockselement "github.com/qkitzero/combination-service/mocks/domain/element"
)

func TestCreateElement(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name             string
		success          bool
		ctx              context.Context
		elementName      string
		createElementErr error
	}{
		{"success create element", true, context.Background(), "test element", nil},
		{"failure create element error", false, context.Background(), "test element", fmt.Errorf("create element error")},
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
			mockCombinationUsecase.EXPECT().CreateElement(tt.elementName).Return(mockElement, tt.createElementErr).AnyTimes()

			combinationHandler := NewCombinationHandler(mockCombinationUsecase)

			req := &combinationv1.CreateElementRequest{
				Name: tt.elementName,
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
