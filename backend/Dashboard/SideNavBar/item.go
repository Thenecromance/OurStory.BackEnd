package SideNavBar

const (
	none = iota
	login
	admin
	developer
)

type item struct {
	RoutePath   string `json:"to"`
	DisplayName string `json:"name"`
	DisplayIcon string `json:"icon"`

	ShowLevel int `json:"-"` // using this to determine if the item should be shown
}
