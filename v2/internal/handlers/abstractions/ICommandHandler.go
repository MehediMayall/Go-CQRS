package abstractions

type ICommandHandler interface {
	Handle() error
}
