package element

type ElementRepository interface {
	Create(element Element) error
}
