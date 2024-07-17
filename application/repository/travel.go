package repository

import (
	"errors"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type HasCreateTravel interface {
	CreateTravel(info *models.Travel) error
}
type HasDeleteTravel interface {
	DeleteTravel(travelId int64) error
}
type HasUpdateTravel interface {
	UpdateTravel(info *models.Travel) error
}
type HasGetTravelByID interface {
	GetTravelByID(travelId int64) (*models.Travel, error)
}
type HasGetTravelByOwner interface {
	GetTravelByOwner(owner int64) ([]models.Travel, error)
}
type HasGetTravelByLocation interface {
	GetTravelByLocation(location string) ([]models.Travel, error)
}
type HasGetTravelByState interface {
	GetTravelByState(state int) ([]models.Travel, error)
}
type HasGetTravelListByID interface {
	GetTravelListByID(id int64) ([]models.Travel, error)
}

type TravelRepository interface {
	Interface.Repository
	CreateTravel(info *models.Travel) error
	DeleteTravel(travelId int64) error
	UpdateTravel(info *models.Travel) error
	GetTravelByID(travelId int64) (*models.Travel, error)
	GetTravelByOwner(owner int64) ([]models.Travel, error)
	GetTravelByLocation(location string) ([]models.Travel, error)
	GetTravelByState(state int) ([]models.Travel, error)
	GetTravelListByID(id int64) ([]models.Travel, error)
}

type travelRepository struct {
	db    *gorp.DbMap
	cache any
}

func (t *travelRepository) BindTable() error {

	t.db.AddTableWithName(models.Travel{}, "Travels")
	t.db.AddTableWithName(models.TravelLog{}, "TravelLogs")
	return nil
}

func (t *travelRepository) CreateTravel(info *models.Travel) error {
	err := t.dbCreateTravel(info)
	if err != nil {
		log.Errorf("CreateTravel error: %v", err)
		return err
	}

	if cache, ok := t.cache.(HasCreateTravel); ok {
		return cache.CreateTravel(info)
	}
	return nil
}

func (t *travelRepository) dbCreateTravel(info *models.Travel) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		_ = trans.Rollback()
		return err
	}

	err = trans.Insert(info)
	if err != nil {
		log.Error(err)
		_ = trans.Rollback()
		return err
	}

	return trans.Commit()
}

func (t *travelRepository) DeleteTravel(travelId int64) error {
	if cache, ok := t.cache.(HasDeleteTravel); ok {
		err := cache.DeleteTravel(travelId)
		if err != nil {
			return err
		} // delete travel data from cache
	}

	return t.dbDeleteTravel(travelId)
}
func (t *travelRepository) dbDeleteTravel(travelId int64) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}
	//delete the travel data from db by id , if not exists
	//errId, err := trans.Delete(models.Travel{}, travelId)
	/*obj, err := t.GetTravelByID(travelId)
	if err != nil {
		return err
	}*/
	errId, err := trans.Query("delete from Travels where travel_id = ?", travelId)
	defer errId.Close()
	if err != nil {
		log.Errorf("DeleteTravel error: %v\n", err)

		return errors.New("delete travel failed")
	}

	return trans.Commit()
}

func (t *travelRepository) UpdateTravel(info *models.Travel) error {
	if cache, ok := t.cache.(HasUpdateTravel); ok {
		err := cache.UpdateTravel(info)
		if err != nil {
			return err

		}
	}

	return t.dbUpdateTravel(info)
}
func (t *travelRepository) dbUpdateTravel(info *models.Travel) error {
	trans, err := t.db.Begin()
	if err != nil {
		log.Error(err)
		return err
	}

	updateId, err := trans.Update(info)
	if err != nil {
		log.Errorf("UpdateTravel error: %v\nerror UserId:%d", err, updateId)
		return err
	}

	return trans.Commit()

}

func (t *travelRepository) GetTravelByID(travelId int64) (*models.Travel, error) {
	if cache, ok := t.cache.(HasGetTravelByID); ok {
		return cache.GetTravelByID(travelId)
	}

	return t.dbGetTravelByID(travelId)

}
func (t *travelRepository) dbGetTravelByID(travelId int64) (*models.Travel, error) {
	//get travel from db by id
	travel := new(models.Travel)
	err := t.db.SelectOne(travel, "select * from Travels where travel_id = ?", travelId)
	if err != nil {
		log.Warnf("GetTravelByID error: %v", err)
		return nil, err
	}

	return travel, nil
}

func (t *travelRepository) GetTravelByOwner(owner int64) ([]models.Travel, error) {
	if cache, ok := t.cache.(HasGetTravelByOwner); ok {
		return cache.GetTravelByOwner(owner)
	}

	return t.dbGetTravelByOwner(owner)
}
func (t *travelRepository) dbGetTravelByOwner(owner int64) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	objects, err := t.db.Select(models.Travel{}, "select * from Travels where user_id = ?", owner)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}

	for _, obj := range objects {
		travel = append(travel, *obj.(*models.Travel))
	}

	return travel, nil
}

func (t *travelRepository) GetTravelByLocation(location string) ([]models.Travel, error) {
	if cache, ok := t.cache.(HasGetTravelByLocation); ok {
		return cache.GetTravelByLocation(location)
	}

	return t.dbGetTravelByLocation(location)
}
func (t *travelRepository) dbGetTravelByLocation(location string) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	err := t.db.SelectOne(travel, "select * from Travels where location = ?", location)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}

	return travel, nil

}

func (t *travelRepository) GetTravelByState(state int) ([]models.Travel, error) {
	if cache, ok := t.cache.(HasGetTravelByState); ok {
		return cache.GetTravelByState(state)
	}

	return t.dbGetTravelByState(state)
}
func (t *travelRepository) dbGetTravelByState(state int) ([]models.Travel, error) {
	//get travel from db by id
	var travel []models.Travel
	err := t.db.SelectOne(travel, "select * from Travels where state = ?", state)
	if err != nil {
		log.Errorf("GetTravelByID error: %v", err)
		return nil, err
	}
	return travel, nil
}

func (t *travelRepository) GetTravelListByID(id int64) ([]models.Travel, error) {
	if cache, ok := t.cache.(HasGetTravelListByID); ok {
		return cache.GetTravelListByID(id)
	}

	return t.dbGetTravelListByID(id)
}
func (t *travelRepository) dbGetTravelListByID(id int64) ([]models.Travel, error) {
	var lists []models.Travel

	objects, err := t.db.Select(models.Travel{}, "select * from Travels where ( user_id = ?) or (find_in_set(?,together) > 0)", id, id)
	if err != nil {
		log.Errorf("GetTravelListByID error: %v", err)
		return nil, err
	}

	for _, obj := range objects {
		lists = append(lists, *obj.(*models.Travel))
	}

	return lists, nil
}

func NewTravelRepository(db *gorp.DbMap) TravelRepository {
	tr := &travelRepository{
		db: db,
	}

	return tr

}
