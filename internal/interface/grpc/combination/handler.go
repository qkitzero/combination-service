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
	element, err := h.combinationUsecase.CreateElement(req.GetName(), req.GetCategoryIds())
	if err != nil {
		return nil, err
	}

	return &combinationv1.CreateElementResponse{
		ElementId: element.ID().String(),
	}, nil
}

func (h *CombinationHandler) ListElements(ctx context.Context, req *combinationv1.ListElementsRequest) (*combinationv1.ListElementsResponse, error) {
	elements, err := h.combinationUsecase.ListElements()
	if err != nil {
		return nil, err
	}

	pbElements := make([]*combinationv1.Element, 0, len(elements))
	for _, element := range elements {
		pbCategories := make([]*combinationv1.Category, 0, len(element.Categories()))
		for _, category := range element.Categories() {
			pbCategory := &combinationv1.Category{
				Id:   category.ID().String(),
				Name: category.Name().String(),
			}
			pbCategories = append(pbCategories, pbCategory)
		}
		pbElement := &combinationv1.Element{
			Id:         element.ID().String(),
			Name:       element.Name().String(),
			Categories: pbCategories,
		}
		pbElements = append(pbElements, pbElement)
	}

	return &combinationv1.ListElementsResponse{
		Elements: pbElements,
	}, nil
}

func (h *CombinationHandler) CreateCategory(ctx context.Context, req *combinationv1.CreateCategoryRequest) (*combinationv1.CreateCategoryResponse, error) {
	category, err := h.combinationUsecase.CreateCategory(req.GetName())
	if err != nil {
		return nil, err
	}

	return &combinationv1.CreateCategoryResponse{
		CategoryId: category.ID().String(),
	}, nil
}

func (h *CombinationHandler) ListCategories(ctx context.Context, req *combinationv1.ListCategoriesRequest) (*combinationv1.ListCategoriesResponse, error) {
	categories, err := h.combinationUsecase.ListCategories()
	if err != nil {
		return nil, err
	}

	pbCategories := make([]*combinationv1.Category, 0, len(categories))
	for _, category := range categories {
		pbCategory := &combinationv1.Category{
			Id:   category.ID().String(),
			Name: category.Name().String(),
		}
		pbCategories = append(pbCategories, pbCategory)
	}

	return &combinationv1.ListCategoriesResponse{
		Categories: pbCategories,
	}, nil
}
