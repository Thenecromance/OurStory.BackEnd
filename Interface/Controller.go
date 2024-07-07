package Interface

type IController interface {
	Initialize()
	Name() string
	SetupRoutes()
	GetRoutes() []IRoute
}
