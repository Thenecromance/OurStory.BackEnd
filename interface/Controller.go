package Interface

type NodeCallback func(path string, parent string, middleWare ...string) *GroupNode
type Controller interface {
	BuildRoutes()

	Name() string

	RequestGroup(NodeCallback)
}
