package abstractions

type IQueryHandler interface {
	Handle() (interface{}, error)
}
