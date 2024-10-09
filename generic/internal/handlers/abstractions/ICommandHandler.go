package abstractions

type ICommandHandler[T any] interface {
	Handle() (T, error)
}
