package repository

import (
	"errors"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

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
	db *gorp.DbMap
}

func (t *travelRepository) BindTable() error {

	t.db.AddTableWithName(models.Travel{}, "Travels")
	t.db.AddTableWithName(models.TravelLog{}, "TravelLogs")
	return nil
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

func (t *travelRepository) DeleteTravel(travelId int64) error {
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
