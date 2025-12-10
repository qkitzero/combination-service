package rule

import (
	"errors"
	"math/rand"
	"time"

	"github.com/qkitzero/combination-service/internal/domain/element"
)

type Strategy interface {
	Combine(elements []element.Element, count int) ([]element.Element, error)
}

type strategy struct {
	r *rand.Rand
}

func NewStrategy() Strategy {
	return &strategy{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (s strategy) Combine(elements []element.Element, count int) ([]element.Element, error) {
	if count < 0 {
		return nil, errors.New("count must not be negative")
	}
	if count > len(elements) {
		return nil, errors.New("count exceeds number of elements")
	}
	if count == 0 {
		return []element.Element{}, nil
	}

	idx := make([]int, len(elements))
	for i := range idx {
		idx[i] = i
	}

	s.r.Shuffle(len(idx), func(i, j int) {
		idx[i], idx[j] = idx[j], idx[i]
	})

	result := make([]element.Element, count)
	for i := 0; i < count; i++ {
		result[i] = elements[idx[i]]
	}

	return result, nil
}
