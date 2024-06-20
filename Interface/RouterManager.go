package Interface

type IRouterController interface {
	// RegisterRouter registers a route to the controller and returns an error if the route is already registered.
	RegisterRouter(routerProxy ...IRoute) error

	// GetRouter returns a route by its name.
	GetRouter(name string) (IRoute, error)

	// ApplyRouter applies all the routers to the gin.Engine.
	ApplyRouter() error

	// Close closes the controller.
	Close() error
}
