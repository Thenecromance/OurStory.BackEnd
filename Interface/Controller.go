package Interface

type IController interface {
	Initialize()
	Name() string
	SetRoutes()
	GetRoutes() []IRoute
}
