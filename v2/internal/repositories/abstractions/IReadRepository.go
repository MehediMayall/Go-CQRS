package repositories

type IReadRepository interface {
	GetAll() (*[]interface{}, error)
	GetById(id string) (*interface{}, error)
}
