package SideNavBar

const (
	none      = iota // when item is set to none, means no matter who you are, you can see this item
	login            // when item is set to login, means only login user can see this item
	admin            //  admin level user can see this item,
	developer        // danger trigger, all the stuff can be seen
)

// Side bar's item object
//
// detail to see the menu item in src\layouts\full\vertical-sidebar\sidebarItem.ts
type item struct {
	Header      string `json:"header"`
	Title       string `json:"title"`
	Icon        string `json:"icon"` // just same as the icon name in src\assets\icons
	To          string `json:"to"`
	GetURL      bool   `json:"getURL"`
	Divider     bool   `json:"divider"`
	Chip        string `json:"chip"`
	ChipColor   string `json:"chipColor"`
	ChipVariant string `json:"chipVariant"`
	ChipIcon    string `json:"chipIcon"`
	Children    []item `json:"children"`
	Disabled    bool   `json:"disabled"`
	TypeString  string `json:"type"`
	SubCaption  string `json:"subCaption"`
}

type itemControl struct {
	VisibleLevel int    `json:"visibleLevel"`
	Items        []item `json:"items"`
}
