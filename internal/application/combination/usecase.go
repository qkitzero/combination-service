package combination

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
)

type CombinationUsecase interface {
	CreateElement(name string) (element.Element, error)
	CreateCategory(name string) (category.Category, error)
}

type combinationUsecase struct {
	elementRepo  element.ElementRepository
	categoryRepo category.CategoryRepository
}

func NewCombinationUsecase(
	elementRepo element.ElementRepository,
	categoryRepo category.CategoryRepository,
) CombinationUsecase {
	return &combinationUsecase{
		elementRepo:  elementRepo,
		categoryRepo: categoryRepo,
	}
}

func (u *combinationUsecase) CreateElement(name string) (element.Element, error) {
	newName, err := element.NewName(name)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newElement := element.NewElement(element.NewElementID(), newName, []category.Category{}, now)

	if err = u.elementRepo.Create(newElement); err != nil {
		return nil, err
	}

	return newElement, nil
}

func (u *combinationUsecase) CreateCategory(name string) (category.Category, error) {
	newName, err := category.NewName(name)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newCategory := category.NewCategory(category.NewCategoryID(), newName, now)

	if err = u.categoryRepo.Create(newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil
}
