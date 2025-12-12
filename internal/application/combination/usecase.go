package combination

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/category"
	"github.com/qkitzero/combination-service/internal/domain/element"
	"github.com/qkitzero/combination-service/internal/domain/language"
	"github.com/qkitzero/combination-service/internal/domain/rule"
)

type CombinationUsecase interface {
	CreateElement(name, languageCode string, categoryIDs []string) (element.Element, error)
	ListElements() ([]element.Element, error)
	CreateCategory(name, languageCode string) (category.Category, error)
	ListCategories() ([]category.Category, error)
	GetCombination(count int) ([]element.Element, error)
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

func (u *combinationUsecase) CreateElement(name, languageCode string, categoryIDs []string) (element.Element, error) {
	newName, err := element.NewName(name)
	if err != nil {
		return nil, err
	}

	newLanguage, err := language.NewLanguage(languageCode)
	if err != nil {
		return nil, err
	}

	cids := make([]category.CategoryID, 0, len(categoryIDs))
	for _, id := range categoryIDs {
		cid, err := category.NewCategoryIDFromString(id)
		if err != nil {
			return nil, err
		}
		cids = append(cids, cid)
	}

	categories, err := u.categoryRepo.FindAllByIDs(cids)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newElement := element.NewElement(element.NewElementID(), newName, newLanguage, categories, now)

	if err = u.elementRepo.Create(newElement); err != nil {
		return nil, err
	}

	return newElement, nil
}

func (u *combinationUsecase) ListElements() ([]element.Element, error) {
	return u.elementRepo.FindAll()
}

func (u *combinationUsecase) CreateCategory(name, languageCode string) (category.Category, error) {
	newName, err := category.NewName(name)
	if err != nil {
		return nil, err
	}

	newLanguage, err := language.NewLanguage(languageCode)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newCategory := category.NewCategory(category.NewCategoryID(), newName, newLanguage, now)

	if err = u.categoryRepo.Create(newCategory); err != nil {
		return nil, err
	}

	return newCategory, nil
}

func (u *combinationUsecase) ListCategories() ([]category.Category, error) {
	return u.categoryRepo.FindAll()
}

func (u *combinationUsecase) GetCombination(count int) ([]element.Element, error) {
	newStrategy, err := rule.NewStrategy(rule.StrategyTypeRandom)
	if err != nil {
		return nil, err
	}

	newRule, err := rule.NewRule(count, newStrategy)
	if err != nil {
		return nil, err
	}

	elements, err := u.elementRepo.FindAll()
	if err != nil {
		return nil, err
	}

	combinationElements, err := newRule.Apply(elements)
	if err != nil {
		return nil, err
	}

	return combinationElements, nil
}
