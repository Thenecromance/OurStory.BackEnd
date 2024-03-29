package Dashboard

import (
	Config "github.com/Thenecromance/OurStories/base/config"
	"github.com/Thenecromance/OurStories/base/logger"
)

type DynamicResource struct {
	Title string `json:"title" ini:"title"`
}

func (d *DynamicResource) load() {
	if err := Config.MapSection("Argon", d); err != nil {
		logger.Get().Errorf("%s faile to map section. ERROR:%s", "Tokens", err)
		return
	}
	if err := Config.ReflectFrom("Argon", d); err != nil {
		logger.Get().Errorf("%s faile to reflect section. ERROR:%s", "Tokens", err)
		return
	}
}

type Icon struct {
	Component  string `json:"component"`
	Background string `json:"background"`
	Shape      string `json:"shape"`
}
type topCardItem struct {
	Title       string `json:"title"`
	Value       string `json:"value"`
	Description string `json:"description"`
	ShowIcon    Icon   `json:"icon"`
}
