package services

import (
	"fmt"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/utility/helper"
	"github.com/Thenecromance/OurStories/utility/log"
	"strconv"
	"strings"
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
	Update(obj *models.Travel) error
}

type travelGetter interface {
	GetTravelByID(id string, userId int) (*models.TravelDTO, error)
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
type travelMap = map[travelId]*models.TravelDTO

// ---------------------------------
// updater object
// ---------------------------------
// using a type to store the travel info , decrease the code duplication
type updatingTravel models.TravelDTO

func (u *updatingTravel) updateState(state int) {
	if state == 0 {
		return
	}
	u.State = state
}
func (u *updatingTravel) updateLocation(location string) {
	if location == "" {
		return
	}
	u.Location = location
}
func (u *updatingTravel) updateDetails(details string) {
	if details == "" {

	}
	u.Details = details
}
func (u *updatingTravel) updateTogetherWith(togetherWith []int) {
	if togetherWith == nil {
		return
	}
	//u.TogetherWith = togetherWith
}
func (u *updatingTravel) updateStartTime(startTime int64) {
	if startTime <= 0 {
		return

	}
	u.StartTime = startTime
}
func (u *updatingTravel) updateEndTime(endTime int64) {
	if endTime <= 0 {
		return
	}
	u.EndTime = endTime
}
func (u *updatingTravel) updateOwner(owner int) {
	if owner <= 0 {
		return
	}
	u.UserId = owner
}

func newUpdater(travel *models.TravelDTO) *updatingTravel {
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

func travelToDTO(travel *models.Travel) *models.TravelDTO {
	dto := &models.TravelDTO{
		Id:           travel.Id,
		State:        travel.State,
		Location:     travel.Location,
		Details:      travel.Details,
		StartTime:    travel.StartTime,
		EndTime:      travel.EndTime,
		TogetherWith: make([]int, 0),
	}
	for _, v := range strings.Split(travel.TogetherWith, ",") {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Warnf("failed convert string to int with error: %v", err)
			continue
		}
		dto.TogetherWith = append(dto.TogetherWith, id)
	}
	return dto
}
func dtoToTravel(travel *models.TravelDTO) *models.Travel {
	obj := &models.Travel{
		Id:           travel.Id,
		State:        travel.State,
		Location:     travel.Location,
		Details:      travel.Details,
		StartTime:    travel.StartTime,
		EndTime:      travel.EndTime,
		TogetherWith: "",
	}
	for _, v := range travel.TogetherWith {
		obj.TogetherWith += strconv.Itoa(v) + ","
	}
	// remove the last comma
	if len(obj.TogetherWith) > 0 {
		obj.TogetherWith = obj.TogetherWith[:len(obj.TogetherWith)-1]
	}
	return obj
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
func (t *travelServiceImpl) userInTravel(travel *models.TravelDTO, userId int) bool {
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

func (t *travelServiceImpl) getTravelData(travelId string) *models.TravelDTO {
	if travel, ok := t.cache[travelId]; ok {
		return travel
	}

	id, err := strconv.Atoi(travelId)
	travel, err := t.repo.GetTravelByID(id)

	if err != nil {
		log.Warnf("getTravelData error: %v", err)
		return nil
	}
	t.cache[travelId] = travelToDTO(travel)

	return travelToDTO(travel)
}

func (t *travelServiceImpl) GetTravelByID(id string, userId int) (*models.TravelDTO, error) {
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

// UpdateToDb update the travel info from cache to db
func (t *travelServiceImpl) UpdateToDb(id string) error {
	log.Debugf("UpdateToDb id: %s", id)

	obj := (*models.TravelDTO)(t.getUpdaterObject(id))
	if obj == nil {
		return fmt.Errorf("UpdateToDb error: %s", "getUpdaterObject return nil")
	}
	// trying to update to db (transaction)
	err := t.repo.UpdateTravel(dtoToTravel(obj))
	if err != nil {
		// if failed , need to handle the cache
		panic("something wrong with db, need to handle the cache here")
		return err
	}

	// if update success, remove the cache, sync back to cache
	t.cache[id] = obj
	delete(t.updatingCache, id) // erase the updating cache
	return nil

}

func (t *travelServiceImpl) Update(obj *models.Travel) error {
	o := t.getUpdaterObject(strconv.Itoa(obj.Id))
	if o == nil {
		return fmt.Errorf("Update error: %s", "getUpdaterObject return nil")
	}
	o.updateDetails(obj.Details)
	o.updateEndTime(obj.EndTime)
	o.updateLocation(obj.Location)
	o.updateOwner(obj.UserId)
	o.updateStartTime(obj.StartTime)
	o.updateState(obj.State)
	//o.updateTogetherWith(obj.TogetherWith)
	err := t.UpdateToDb(strconv.Itoa(o.Id))
	if err != nil {
		log.Errorf("Update error: %v", err)
		return err
	}
	return nil

}

func (t *travelServiceImpl) CreateTravel(info *models.Travel) error {
	if t.repo == nil {
		panic("repo is nil")
	}
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
