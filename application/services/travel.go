package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"
)

//---------------------------------
// TravelService interface define
//---------------------------------

type travelUpdater interface {
	UpdateState(id string, state int) error
	UpdateLocation(id string, location string) error
	UpdateDetails(id string, details string) error
	UpdateTogetherWith(id string, togetherWith []int64) error
	UpdateStartTime(id string, startTime int64) error
	UpdateEndTime(id string, endTime int64) error
	UpdateOwner(id string, owner int64) error
	UpdateToDb(id string) error
	Update(obj *models.Travel) error
}

type travelGetter interface {
	GetTravelByID(id string, userId int64) (*models.Travel, error)
	GetTravelByOwner(owner int64) ([]models.Travel, error)
	GetTravelByLocation(location string) ([]models.Travel, error)
	GetTravelByState(state int) ([]models.Travel, error)
	GetTravelList(userId int64) ([]models.Travel, error)
}

// TravelService is interface for travel service
type TravelService interface {
	travelGetter
	travelUpdater

	CreateTravel(dto *models.Travel) error
	DeleteTravel(id string, userId int64) error
}

type travelId = string

// todo: use ICache to store the travel info,due to ICache only implements LRU, so  temp to use map instead.
type travelMap = map[travelId]*models.Travel

// ---------------------------------
// updater object
// ---------------------------------
// using a type to store the travel info , decrease the code duplication
type updatingTravel struct {
	*models.Travel
	needUpdate bool
}

func (u *updatingTravel) updateState(state int) {
	if state == 0 {
		return
	}
	u.needUpdate = true
	u.State = state
}
func (u *updatingTravel) updateLocation(location string) {
	if location == "" {
		return
	}
	u.needUpdate = true
	u.Location = location
}
func (u *updatingTravel) updateDetails(details string) {
	if details == "" {
		return
	}
	u.needUpdate = true
	u.Detail = details
}
func (u *updatingTravel) updateTogetherWith(togetherWith []int64) {
	if togetherWith == nil {
		return
	}
	u.needUpdate = true
	//u.TogetherWith = togetherWith
}
func (u *updatingTravel) updateStartTime(startTime int64) {
	if startTime <= 0 {
		return

	}
	u.needUpdate = true
	u.StartTime = startTime
}
func (u *updatingTravel) updateEndTime(endTime int64) {
	if endTime <= 0 {
		return
	}
	u.needUpdate = true
	u.EndTime = endTime
}
func (u *updatingTravel) updateOwner(owner int64) {
	if owner <= 0 {
		return
	}
	u.needUpdate = true
	u.UserId = owner
}

func newUpdater(travel *models.Travel) *updatingTravel {
	return &updatingTravel{
		Travel:     travel,
		needUpdate: false,
	}

}

//---------------------------------
// TravelService implement
//---------------------------------

type travelServiceImpl struct {
	repo          repository.TravelRepository
	cache         travelMap
	updatingCache map[travelId]*updatingTravel
}

func parseTravelState(travel *models.Travel) {
	now := time.Now().Unix()
	if travel.StartTime > now {
		travel.State = models.TravelStatePending
	} else if travel.EndTime < now {
		travel.State = models.TravelStateFinished
	} else {
		travel.State = models.TravelStateOngoing
	}
}

func (t *travelServiceImpl) getUpdaterObject(id string) *updatingTravel {
	travel, exists := t.updatingCache[id]
	if !exists {
		travelInDb, err := t.GetTravelByID(id, 0)
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

// userInTravel check if the user in the travel
func (t *travelServiceImpl) userInTravel(travel *models.Travel, userId int64) bool {
	if userId == travel.UserId {
		return true
	}
	for _, v := range travel.TogetherWith {
		if v == userId {
			return true
		}
	}
	return false
}

func (t *travelServiceImpl) getTravelData(travelId string) *models.Travel {
	if travel, ok := t.cache[travelId]; ok {
		return travel
	}

	//format the string to int64
	id, err := strconv.ParseInt(travelId, 10, 64)
	if err != nil {
		log.Warnf("getTravelData error: %v", err)
		return nil
	}

	travel, err := t.repo.GetTravelByID(id)

	if err != nil {
		log.Warnf("getTravelData error: %v", err)
		return nil
	}
	t.cache[travelId] = travel

	return travel
}

func (t *travelServiceImpl) GetTravelByID(id string, userId int64) (*models.Travel, error) {
	log.Debug("start to get travel by id:", id)
	dto := t.getTravelData(id)
	if dto == nil {
		return nil, fmt.Errorf("could not find travel id %s", id)
	}

	log.Debug("check user in travel...")
	if t.userInTravel(dto, userId) {

		return dto, nil
	}

	log.Debug("user not in travel")
	return nil, fmt.Errorf("user not in travel")
}

func (t *travelServiceImpl) GetTravelByOwner(owner int64) ([]models.Travel, error) {
	travels, err := t.repo.GetTravelByOwner(owner)
	if err != nil {
		log.Errorf("GetTravelByOwner error: %v", err)
		return nil, errors.New("no travel data in db")
	}
	var returnObjs []models.Travel
	for _, v := range travels {
		travel := v
		returnObjs = append(returnObjs, travel)
	}
	return returnObjs, nil
}

func (t *travelServiceImpl) GetTravelByLocation(location string) ([]models.Travel, error) {
	travels, err := t.repo.GetTravelByLocation(location)
	if err != nil {
		log.Errorf("GetTravelByLocation error: %v", err)
		return nil, errors.New("no travel data in db")
	}

	var returnObjs []models.Travel
	for _, v := range travels {
		travel := v
		returnObjs = append(returnObjs, travel)
	}

	return returnObjs, nil
}

func (t *travelServiceImpl) GetTravelByState(state int) ([]models.Travel, error) {
	travels, err := t.repo.GetTravelByState(state)
	if err != nil {
		log.Errorf("GetTravelByState error: %v", err)
		return nil, errors.New("no travel data in db")
	}

	var returnObjs []models.Travel
	for _, v := range travels {
		travel := v
		returnObjs = append(returnObjs, travel)
	}

	return returnObjs, nil
}

func (t *travelServiceImpl) GetTravelList(userId int64) ([]models.Travel, error) {
	travels, err := t.repo.GetTravelListByID(userId)
	if err != nil {
		log.Errorf("GetTravelList error: %v", err)
		return nil, errors.New("no travel data in db")
	}

	var returnObjs []models.Travel
	for _, v := range travels {
		travel := v
		returnObjs = append(returnObjs, travel)
	}

	return returnObjs, nil
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

func (t *travelServiceImpl) UpdateTogetherWith(id string, togetherWith []int64) error {
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

func (t *travelServiceImpl) UpdateOwner(id string, owner int64) error {
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

// UpdateToDb update the travel info from cache to db
func (t *travelServiceImpl) UpdateToDb(id string) error {
	log.Debugf("UpdateToDb id: %s", id)

	obj := t.getUpdaterObject(id)
	if obj == nil {
		return fmt.Errorf("UpdateToDb error: %s", "getUpdaterObject return nil")
	}
	// trying to update to db (transaction)
	err := t.repo.UpdateTravel(obj.Travel)
	if err != nil {
		// if failed , need to handle the cache
		panic("something wrong with db, need to handle the cache here")
		return err
	}

	// if update success, remove the cache, sync back to cache
	t.cache[id] = obj.Travel
	delete(t.updatingCache, id) // erase the updating cache
	return nil

}

func (t *travelServiceImpl) Update(obj *models.Travel) error {

	o := t.getUpdaterObject(strconv.FormatInt(obj.Id, 10))
	if o == nil {
		return fmt.Errorf("update error: %s", "getUpdaterObject return nil")
	}
	o.updateDetails(obj.Detail)
	o.updateEndTime(obj.EndTime)
	o.updateLocation(obj.Location)
	o.updateOwner(obj.UserId)
	o.updateStartTime(obj.StartTime)
	o.updateState(obj.State)
	o.updateTogetherWith(obj.TogetherWith)

	// if there is no change, no need to update
	if !o.needUpdate {
		return nil
	}

	err := t.UpdateToDb(strconv.FormatInt(o.Id, 10))
	if err != nil {
		log.Errorf("Update error: %v", err)
		return err
	}
	return nil

}

func (t *travelServiceImpl) CreateTravel(travel *models.Travel) error {
	if t.repo == nil {
		panic("repo is nil")
	}

	//travel := dtoToTravel(dto)
	return t.repo.CreateTravel(travel)
}

func (t *travelServiceImpl) DeleteTravel(id string, userId int64) error {

	dto, err := t.GetTravelByID(id, userId)
	if err != nil {
		log.Errorf("DeleteTravel error: %v", err)
		return fmt.Errorf("could not find travel id %s", id)
	}
	if dto == nil {
		return fmt.Errorf("could not find travel id %s", id)
	}

	iid, err := strconv.ParseInt(id, 10, 64)
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
