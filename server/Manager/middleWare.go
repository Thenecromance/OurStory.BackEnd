package Manager

import "github.com/Thenecromance/OurStories/Interface"

type MiddleWareManager struct {
}

func NewMiddleWareManager() Interface.MiddleWareController {
	return &MiddleWareManager{}
}
