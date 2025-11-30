package combination

import (
	"context"

	combinationv1 "github.com/qkitzero/combination-service/gen/go/combination/v1"
	appcombination "github.com/qkitzero/combination-service/internal/application/combination"
)

type CombinationHandler struct {
	combinationv1.UnimplementedCombinationServiceServer
	combinationUsecase appcombination.CombinationUsecase
}

func NewCombinationHandler(combinationUsecase appcombination.CombinationUsecase) *CombinationHandler {
	return &CombinationHandler{
		combinationUsecase: combinationUsecase,
	}
}

func (h *CombinationHandler) CreateElement(ctx context.Context, req *combinationv1.CreateElementRequest) (*combinationv1.CreateElementResponse, error) {
	element, err := h.combinationUsecase.CreateElement(req.GetName())
	if err != nil {
		return nil, err
	}

	return &combinationv1.CreateElementResponse{
		ElementId: element.ID().String(),
	}, nil
}
