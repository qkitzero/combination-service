package rule

import "errors"

var (
	ErrInvalidCount        = errors.New("invalid count")
	ErrUnknownStrategyType = errors.New("unknown strategy type")
)
