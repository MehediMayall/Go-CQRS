package repositories

type IWriteRepository interface {
	Add(interface{}) error
	Update(interface{}) error
	Delete(id string) error
}
