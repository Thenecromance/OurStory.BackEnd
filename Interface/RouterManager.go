package Interface

type RouterController interface {
	// RegisterRouter registers a route to the controller and returns an error if the route is already registered.
	RegisterRouter(routerProxy ...Route) error

	// GetRouter returns a route by its name.
	GetRouter(name string) (Route, error)

	// ApplyRouter applies all the routers to the gin.Engine.
	ApplyRouter() error

	// Close closes the controller.
	Close() error
}
