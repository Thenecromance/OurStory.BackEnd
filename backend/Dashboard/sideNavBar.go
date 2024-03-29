package Dashboard

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/base/logger"
	"os"
)

type sideNavBar struct {
	To      string `json:"to"`
	NavText string `json:"navText"`
	Icon    string `json:"icon"`
	Msg     string `json:"msg"`
	IsSplit bool   `json:"isSplit"`
}

type SideNavBarModel struct {
	Navs []sideNavBar `json:"navs"`
}

func (sb *SideNavBarModel) defaultConfig() {
	sb.Navs = append(sb.Navs, sideNavBar{
		To:      "/",
		NavText: "Dashboard",
		Icon:    "ni-tv-2 text-primary",
		Msg:     "",
		IsSplit: false,
	})
}

func (sb *SideNavBarModel) Load() {
	sb.Navs = make([]sideNavBar, 0)

	_, err := os.Stat("setting/sideNavBar.json")
	if err != nil {
		sb.defaultConfig()
		marshal, err := json.Marshal(sb.Navs)
		if err != nil {
			return
		}
		os.WriteFile("setting/sideNavBar.json", marshal, 0644)
		return
	}
	file, err := os.ReadFile("setting/sideNavBar.json")
	if err != nil {
		logger.Get().Info("fail to read sideNavBar.json", err)
		return
	}

	err = json.Unmarshal(file, &sb.Navs)
	if err != nil {
		logger.Get().Info("fail to unmarshal sideNavBar.json", err)
		return
	}

}
