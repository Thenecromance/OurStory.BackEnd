package Interface

type RouterController interface {
	// RegisterRouter registers a router to the controller and returns an error if the router is already registered.
	RegisterRouter(routerProxy Router) error

	// GetRouter returns a router by its name.
	GetRouter(name string) (Router, error)

	// ApplyRouter applies all the routers to the gin.Engine.
	ApplyRouter() error

	// Close closes the controller.
	Close() error
}
