package services

import (
	"fmt"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"
	"strconv"
)

//---------------------------------
// TravelService interface define
//---------------------------------

type travelUpdater interface {
	UpdateState(id string, state int) error
	UpdateLocation(id string, location string) error
	UpdateDetails(id string, details string) error
	UpdateTogetherWith(id string, togetherWith []int) error
	UpdateStartTime(id string, startTime int64) error
	UpdateEndTime(id string, endTime int64) error
	UpdateOwner(id string, owner int) error
	UpdateToDb(id string) error
}

type travelGetter interface {
	GetTravelByID(id string) (*models.Travel, error)
	GetTravelByOwner(owner int) ([]models.Travel, error)
	GetTravelByLocation(location string) ([]models.Travel, error)
	GetTravelByState(state int) ([]models.Travel, error)
}

// TravelService is interface for travel service
type TravelService interface {
	travelGetter
	travelUpdater

	CreateTravel(*models.Travel) error
	DeleteTravel(id string) error
}

type travelId = string

// todo: use ICache to store the travel info,due to ICache only implements LRU, so  temp to use map instead.
type travelMap = map[travelId]*models.Travel

// ---------------------------------
// updater object
// ---------------------------------
// using a type to store the travel info , decrease the code duplication
type updatingTravel models.Travel

func (u *updatingTravel) updateState(state int) {
	u.State = state
}
func (u *updatingTravel) updateLocation(location string) {
	u.Location = location
}
func (u *updatingTravel) updateDetails(details string) {
	u.Details = details
}
func (u *updatingTravel) updateTogetherWith(togetherWith []int) {
	u.TogetherWith = togetherWith
}
func (u *updatingTravel) updateStartTime(startTime int64) {
	u.StartTime = startTime
}
func (u *updatingTravel) updateEndTime(endTime int64) {
	u.EndTime = endTime
}
func (u *updatingTravel) updateOwner(owner int) {
	u.UserId = owner
}

func newUpdater(travel *models.Travel) *updatingTravel {
	return (*updatingTravel)(travel)
}

//---------------------------------
// TravelService implement
//---------------------------------

type travelServiceImpl struct {
	repo          repository.TravelRepository
	cache         travelMap
	updatingCache map[travelId]*updatingTravel
}

func (t *travelServiceImpl) getUpdaterObject(id string) *updatingTravel {
	travel, exists := t.updatingCache[id]
	if !exists {
		travelInDb, err := t.GetTravelByID(id)
		if err != nil {
			log.Errorf("%s  error: %v", helper.GetFunctionName(t.UpdateState), err)
			return nil
		}
		obj := newUpdater(travelInDb)
		t.updatingCache[id] = obj
		travel = t.updatingCache[id]
	}

	return travel
}

func (t *travelServiceImpl) GetTravelByID(id string) (*models.Travel, error) {
	if travel, ok := t.cache[id]; ok {
		return travel, nil
	}

	// get travel from db
	Iid, err := strconv.Atoi(id)
	travel, err := t.repo.GetTravelByID(Iid)
	if err != nil {
		log.Error("GetTravelByID error: %v", err)
		return nil, err
	}

	t.cache[id] = travel
	return travel, nil
}

func (t *travelServiceImpl) GetTravelByOwner(owner int) ([]models.Travel, error) {
	return t.repo.GetTravelByOwner(owner)
}

func (t *travelServiceImpl) GetTravelByLocation(location string) ([]models.Travel, error) {
	return t.repo.GetTravelByLocation(location)
}

func (t *travelServiceImpl) GetTravelByState(state int) ([]models.Travel, error) {
	return t.repo.GetTravelByState(state)
}

func (t *travelServiceImpl) UpdateState(id string, state int) error {
	log.Debugf("UpdateState id: %s, state: %d", id, state)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateState error: %s", "getUpdaterObject return nil")
	}
	obj.updateState(state)
	return nil
}

func (t *travelServiceImpl) UpdateLocation(id string, location string) error {
	log.Debugf("UpdateLocation id: %s, location: %s", id, location)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateLocation error: %s", "getUpdaterObject return nil")
	}
	obj.updateLocation(location)
	return nil
}

func (t *travelServiceImpl) UpdateDetails(id string, details string) error {
	log.Debugf("UpdateDetails id: %s, details: %s", id, details)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateDetails error: %s", "getUpdaterObject return nil")

	}
	obj.updateDetails(details)
	return nil
}

func (t *travelServiceImpl) UpdateTogetherWith(id string, togetherWith []int) error {
	log.Debugf("UpdateTogetherWith id: %s, togetherWith: %v", id, togetherWith)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateTogetherWith error: %s", "getUpdaterObject return nil")
	}
	obj.updateTogetherWith(togetherWith)
	return nil
}

func (t *travelServiceImpl) UpdateStartTime(id string, startTime int64) error {
	log.Debugf("UpdateStartTime id: %s, startTime: %d", id, startTime)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateStartTime error: %s", "getUpdaterObject return nil")

	}
	obj.updateStartTime(startTime)
	return nil
}

func (t *travelServiceImpl) UpdateEndTime(id string, endTime int64) error {
	log.Debugf("UpdateEndTime id: %s, endTime: %d", id, endTime)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateEndTime error: %s", "getUpdaterObject return nil")
	}
	obj.updateEndTime(endTime)
	return nil
}

func (t *travelServiceImpl) UpdateOwner(id string, owner int) error {
	/*travel, err := t.GetTravelByID(id)
	if err != nil {
		return err
	}

	travel.UserId = owner
	t.cache[id] = travel // sync back to cache
	return t.repo.UpdateTravel(travel)*/
	log.Debugf("UpdateOwner id: %s, owner: %d", id, owner)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateOwner error: %s", "getUpdaterObject return nil")
	}
	obj.updateOwner(owner)
	return nil
}

func (t *travelServiceImpl) UpdateToDb(id string) error {
	log.Debugf("UpdateToDb id: %s", id)
	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateToDb error: %s", "getUpdaterObject return nil")
	}
	// trying to update to db (transaction)
	err := t.repo.UpdateTravel((*models.Travel)(obj))
	if err != nil {
		// if failed , need to handle the cache
		panic("something wrong with db, need to handle the cache here")
		return err
	}
	// if update success, remove the cache, sync back to cache
	travel := (*models.Travel)(obj)
	t.cache[id] = travel
	delete(t.updatingCache, id) // erase the updating cache
	return nil

}

func (t *travelServiceImpl) CreateTravel(info *models.Travel) error {

	return t.repo.CreateTravel(info)
}

func (t *travelServiceImpl) DeleteTravel(id string) error {
	/*	if err := t.repo.DeleteTravel(id); err != nil {

		}*/
	iid, err := strconv.Atoi(id)
	if err != nil {
		return err

	}
	err = t.repo.DeleteTravel(iid)
	if err != nil {
		return err
	}
	//remove the cache if exists
	if _, ok := t.cache[id]; ok {
		delete(t.cache, id)
	}
	if _, ok := t.updatingCache[id]; ok {
		log.Error("DeleteTravel error: %s", "updatingCache should not be contains the id after delete it from db")
		delete(t.updatingCache, id)
	}
	return nil
}

func NewTravelService(repo repository.TravelRepository) TravelService {
	return &travelServiceImpl{
		repo,
		make(travelMap),
		make(map[travelId]*updatingTravel),
	}
}
