package score

import (
	"fmt"
)

type Like int

func (l Like) Int() int {
	return int(l)
}

func (l Like) Increment() Like {
	return l + 1
}

func NewLike(value int) (Like, error) {
	if value < 0 {
		return 0, fmt.Errorf("like must be positive")
	}
	return Like(value), nil
}
