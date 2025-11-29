package combination

type CombinationUsecase interface {
	CreateElement() error
}

type combinationUsecase struct {
}

func NewCombinationUsecase() CombinationUsecase {
	return &combinationUsecase{}
}

func (s *combinationUsecase) CreateElement() error {
	return nil
}
