package category

type CategoryRepository interface {
	Create(category Category) error
}
