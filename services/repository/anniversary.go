package repository

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type Anniversary interface {
	CreateAnniversary(anniversary *models.Anniversary) error
	RemoveAnniversary(anniversary *models.Anniversary) error
	UpdateAnniversary(anniversary *models.Anniversary) error
	GetAnniversaryById(id int64) (*models.Anniversary, error)
	GetAnniversaryList(user string) ([]models.Anniversary, error)
}

type anniversaryRepository struct {
	db *gorp.DbMap
}

func (ar *anniversaryRepository) initTable() {
	if ar.db == nil {
		log.Error("db is nil")
		return
	}

	log.Info("start to binding anniversary with table travel")
	tbl := ar.db.AddTableWithName(models.Anniversary{}, "anniversary")
	tbl.SetKeys(false, "Id") // using snowflake to generate the id
	tbl.ColMap("Id").SetNotNull(true)
	tbl.ColMap("Owner").SetNotNull(true)
	tbl.ColMap("TimeStamp").SetNotNull(true)
	tbl.ColMap("Title").SetNotNull(true)
	tbl.ColMap("Info").SetNotNull(true)

	err := ar.db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table anniversary with error: %s", err.Error())
		return
	}
}

//func NewAnniversaryRepository(db *gorp.DbMap) Anniversary {
//	return &anniversaryRepository{db}
//}
