package rule

import (
	"math/rand"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type randomStrategy struct{}

func (s *randomStrategy) Select(count int, elements []element.Element) ([]element.Element, error) {
	indices := rand.Perm(len(elements))

	result := make([]element.Element, count)
	for i := range count {
		result[i] = elements[indices[i]]
	}

	return result, nil
}

func newRandomStrategy() Strategy {
	return &randomStrategy{}
}
