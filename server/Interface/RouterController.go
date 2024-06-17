package Interface

type RouterController interface {
	//Register a router to the server
	RegisterRouter(routerProxy RouterProxy) error

	Close() error
}
