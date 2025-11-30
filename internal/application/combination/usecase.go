package combination

import (
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type CombinationUsecase interface {
	CreateElement(name string) (element.Element, error)
}

type combinationUsecase struct {
	repo element.ElementRepository
}

func NewCombinationUsecase(repo element.ElementRepository) CombinationUsecase {
	return &combinationUsecase{repo: repo}
}

func (u *combinationUsecase) CreateElement(name string) (element.Element, error) {
	newName, err := element.NewName(name)
	if err != nil {
		return nil, err
	}

	now := time.Now()

	newElement := element.NewElement(element.NewElementID(), newName, now)

	if err = u.repo.Create(newElement); err != nil {
		return nil, err
	}

	return newElement, nil
}
