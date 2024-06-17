package demo

import "github.com/Thenecromance/OurStories/server/Interface"

type ServiceDemo struct {
	ModelDemo
	ViewDemo
	ControllerDemo
}

type ControllerDemo struct {
	Router Interface.RouterProxy
}

type ModelDemo struct {
}

type ViewDemo struct {
}
