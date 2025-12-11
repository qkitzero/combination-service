package rule

import (
	"github.com/qkitzero/combination-service/internal/domain/element"
)

type StrategyType string

const (
	StrategyTypeRandom StrategyType = "random"
)

type Strategy interface {
	Select(count int, elements []element.Element) ([]element.Element, error)
}

func NewStrategy(strategyType StrategyType) (Strategy, error) {
	switch strategyType {
	case StrategyTypeRandom:
		return newRandomStrategy(), nil
	default:
		return nil, ErrUnknownStrategyType
	}
}
