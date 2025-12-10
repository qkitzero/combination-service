package rule

import (
	"math/rand"
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type randomStrategy struct {
	r *rand.Rand
}

func (s *randomStrategy) Select(count int, elements []element.Element) ([]element.Element, error) {
	indices := s.r.Perm(len(elements))

	result := make([]element.Element, count)
	for i := range count {
		result[i] = elements[indices[i]]
	}

	return result, nil
}

func newRandomStrategy() Strategy {
	return &randomStrategy{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
