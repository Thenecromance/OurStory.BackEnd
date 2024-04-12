package SideNavBar

import (
	"github.com/Thenecromance/OurStories/base/fileWatcher"
	"github.com/Thenecromance/OurStories/base/logger"
	"github.com/Thenecromance/OurStories/base/utils"
)

const (
	file = "setting/side_nav_bar.json"
)

func New() *Model {
	model := &Model{}
	model.initialize()

	return model
}

type Model struct {
	items []itemControl
}

func (m *Model) initialize() {
	m.loadNavItems()
	fileWatcher.WatchFile(file, fileWatcher.FileCallback{
		OnChanged: m.loadNavItems,
	})
}

func (m *Model) loadNavItems() {
	if !utils.FileExists(file) {
		logger.Get().Infof("%s not found, creating default", file)
		m.demoItem()
		utils.SaveJson(file, m.items)
		return
	}
	err := utils.LoadJson(file, &m.items)
	if err != nil {
		logger.Get().Errorf("failed to load %s: %s", file, err)
		return
	}
	logger.Get().Infof("loaded %s", file)
	return
}

func (m *Model) demoItem() {
	m.items = append(m.items, itemControl{
		VisibleLevel: none,
		Items: []item{
			{
				Header: "Dashboard",
				Title:  "Default",
				Icon:   "BeachIcon",
				To:     "/dashboard/default",
			},
		},
	})
}

func (m *Model) Items(lvl int) (list []item) {
	list = make([]item, 0)
	for _, v := range m.items {
		if v.VisibleLevel == lvl {
			list = append(list, v.Items...)
		}
	}
	return
}
