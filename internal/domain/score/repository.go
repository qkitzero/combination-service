package score

type ScoreRepository interface {
	Create(score Score) error
	Update(score Score) error
}
