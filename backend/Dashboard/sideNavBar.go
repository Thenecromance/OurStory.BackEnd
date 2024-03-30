package Dashboard

import (
	"encoding/json"
	"github.com/Thenecromance/OurStories/base/fileWatcher"
	"github.com/Thenecromance/OurStories/base/logger"
	"os"
)

const (
	configPath = "setting/sideNavBar.json"
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

func (sb *SideNavBarModel) reload() {
	sb.Navs = make([]sideNavBar, 0)

	_, err := os.Stat(configPath)
	if err != nil {
		sb.defaultConfig()
		marshal, err := json.Marshal(sb.Navs)
		if err != nil {
			return
		}
		os.WriteFile(configPath, marshal, 0644)
		return
	}
	file, err := os.ReadFile(configPath)
	if err != nil {
		logger.Get().Info("fail to read sideNavBar.json ", err)
		return
	}

	err = json.Unmarshal(file, &sb.Navs)
	if err != nil {
		logger.Get().Info("fail to unmarshal sideNavBar.json ", err)
		return
	}
}

func (sb *SideNavBarModel) Load() {
	sb.reload()

	fileWatcher.WatchFile(configPath, fileWatcher.FileCallback{
		OnChanged: sb.reload,
	})

}
