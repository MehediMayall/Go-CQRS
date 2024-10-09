package repositories

type IReadRepository[T any] interface {
	GetAll() (*[]T, error)
	GetById(id string) (*T, error)
}
