package category

type CategoryRepository interface {
	Create(category Category) error
	FindByID(id CategoryID) (Category, error)
}
