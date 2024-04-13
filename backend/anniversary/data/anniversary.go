package data

import (
	"github.com/Thenecromance/OurStories/base/log"
	"gopkg.in/gorp.v2"
)

type Anniversary struct {
	Id        string `json:"id" db:"id"`
	Owner     string `json:"owner" db:"owner"`
	TimeStamp int64  `json:"time_stamp" db:"time_stamp"`
	Title     string `json:"title" db:"title"`
	Info      string `json:"info" db:"info"`
}

func (a Anniversary) SetupTable(db *gorp.DbMap) {
	log.Info("start to binding anniversary with table travel")
	tbl := db.AddTableWithName(a, "anniversary")
	tbl.SetKeys(false, "Id") // using snowflake to generate the id
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("Owner").SetNotNull(true)
	tbl.ColMap("TimeStamp").SetNotNull(true)
	tbl.ColMap("Title").SetNotNull(true)
	tbl.ColMap("Info").SetNotNull(true)

	err := db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table anniversary with error: %s", err.Error())
		return
	}
}
