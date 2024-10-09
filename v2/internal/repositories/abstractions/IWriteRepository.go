package repositories

type IWriteRepository[T any] interface {
	Add(*T) error
	Update(*T) error
	Delete(id string) error
}
