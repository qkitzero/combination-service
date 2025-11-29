package combination

import (
	"context"
	"fmt"

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
	fmt.Println(req.GetName())

	return &combinationv1.CreateElementResponse{
		ElementId: "element id",
	}, nil
}
