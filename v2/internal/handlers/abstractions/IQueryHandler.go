package abstractions

type IQueryHandle interface {
	Handle() (interface{}, error)
}
