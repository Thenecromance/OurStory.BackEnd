package SideNavBar

const (
	file = "setting/side_nav_bar.json"
)

func New() *Controller {
	ctrl := &Controller{
		items: []item{},
	}

	return ctrl
}

type Controller struct {
	items []item
}

func (c *Controller) initialize() {

}
