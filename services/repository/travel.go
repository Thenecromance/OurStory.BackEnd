package repository

import (
	"fmt"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type TravelRepository interface {
	CreateTravel(info *models.Travel) error
	DeleteTravel(travelId int) error
	UpdateTravel(info *models.Travel) error
	GetTravelByID(travelId int) (*models.Travel, error)
	GetTravelByOwner(owner int) ([]models.Travel, error)
	GetTravelByLocation(location string) ([]models.Travel, error)
	GetTravelByState(state int) ([]models.Travel, error)
}

type travelRepository struct {
	db *gorp.DbMap
}

func (t *travelRepository) CreateTravel(info *models.Travel) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		trans.Rollback()
		return err
	}

	err = trans.Insert(info)
	if err != nil {
		log.Error(err)
		trans.Rollback()
		return err
	}

	return trans.Commit()
}

func (t *travelRepository) DeleteTravel(travelId int) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}
	//delete the travel data from db by id , if not exists
	errId, err := trans.Delete(models.Travel{}, travelId)
	if err != nil {
		log.Errorf("DeleteTravel error: %v\nerror Id:%d", err, errId)
		return err
	}

	return trans.Commit()
}

func (t *travelRepository) UpdateTravel(info *models.Travel) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}

	updateId, err := trans.Update(info)
	if err != nil {
		log.Errorf("UpdateTravel error: %v\nerror Id:%d", err, updateId)
		return err
	}

	return trans.Commit()

}

func (t *travelRepository) GetTravelByID(travelId int) (*models.Travel, error) {
	//get travel from db by id
	travel := new(models.Travel)
	err := t.db.SelectOne(travel, "select * from travel where id = ?", travelId)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}

	return travel, nil
}

func (t *travelRepository) GetTravelByOwner(owner int) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	err := t.db.SelectOne(travel, "select * from travel where owner = ?", owner)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}

	return travel, nil
}

func (t *travelRepository) GetTravelByLocation(location string) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	err := t.db.SelectOne(travel, "select * from travel where location = ?", location)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}

	return travel, nil

}

func (t *travelRepository) GetTravelByState(state int) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	err := t.db.SelectOne(travel, "select * from travel where state = ?", state)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}
	return travel, nil
}

func (t *travelRepository) initTable() error {
	if t.db == nil {
		log.Debugf("db is nil")
		return fmt.Errorf("db is nil")
	}

	log.Infof("start to binding user with table user")
	tbl := t.db.AddTableWithName(models.Travel{}, "travel")
	tbl.SetKeys(true, "Id") // using snowflake to generate the id

	err := t.db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table user with error: %s", err.Error())
		return err
	}
	return nil
}

func NewTravelRepository(db *gorp.DbMap) TravelRepository {
	tr := &travelRepository{
		db: db,
	}
	err := tr.initTable()
	if err != nil {
		panic(err)
		return nil
	}

	return tr

}
