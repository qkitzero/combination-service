package rule

import (
	"math/rand"
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type Strategy interface {
	Random(elements []element.Element, count int) ([]element.Element, error)
}

type strategy struct {
	r *rand.Rand
}

func (s strategy) Random(elements []element.Element, count int) ([]element.Element, error) {
	idx := make([]int, len(elements))
	for i := range idx {
		idx[i] = i
	}

	s.r.Shuffle(len(idx), func(i, j int) { idx[i], idx[j] = idx[j], idx[i] })

	result := make([]element.Element, count)
	for i := range count {
		result[i] = elements[idx[i]]
	}

	return result, nil
}

func NewStrategy() Strategy {
	return &strategy{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}
