package rule

import (
	"github.com/qkitzero/combination-service/internal/domain/element"
)

type Rule interface {
	Count() int
	Strategy() Strategy
	Apply(elements []element.Element) ([]element.Element, error)
}

type rule struct {
	count    int
	strategy Strategy
}

func (r rule) Count() int {
	return r.count
}

func (r rule) Strategy() Strategy {
	return r.strategy
}

func (r rule) Apply(elements []element.Element) ([]element.Element, error) {
	if r.count < 0 {
		return nil, ErrInvalidCount
	}

	if r.count > len(elements) {
		return nil, ErrInvalidCount
	}

	if r.count == 0 {
		return []element.Element{}, nil
	}

	return r.strategy.Select(r.count, elements)
}

func NewRule(
	count int,
	strategy Strategy,
) Rule {
	return &rule{
		count:    count,
		strategy: strategy,
	}
}
