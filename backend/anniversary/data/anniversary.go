package data

import (
	"github.com/Thenecromance/OurStories/base/logger"
	"gopkg.in/gorp.v2"
)

type Anniversary struct {
	Id    string `json:"id" db:"id"`
	Owner string `json:"owner" db:"owner"`
	Year  int    `json:"year" db:"year"`
	Month int    `json:"month" db:"month"`
	Day   int    `json:"day" db:"day"`
	Title string `json:"title" db:"title"`
	Info  string `json:"info" db:"info"`
}

func (a Anniversary) SetupTable(db *gorp.DbMap) {
	logger.Get().Info("start to binding anniversary with table travel")
	tbl := db.AddTableWithName(a, "travel")
	tbl.SetKeys(false, "Id") // using snowflake to generate the id
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("Owner").SetNotNull(true)
	tbl.ColMap("Year").SetNotNull(true)
	tbl.ColMap("Month").SetNotNull(true)
	tbl.ColMap("Day").SetNotNull(true)
	tbl.ColMap("Title").SetNotNull(true)
	tbl.ColMap("Info").SetNotNull(true)

	err := db.CreateTablesIfNotExists()
	if err != nil {
		logger.Get().Errorf("failed to create table anniversary with error: %s", err.Error())
		return
	}
}
