package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type travelUpdater interface {
	UpdateState(id string, state int) error
	UpdateLocation(id string, location string) error
	UpdateDetails(id string, details string) error
	UpdateTogetherWith(id string, togetherWith []int) error
	UpdateStartTime(id string, startTime int64) error
	UpdateEndTime(id string, endTime int64) error
	UpdateOwner(id string, owner int) error
}

type travelGetter interface {
	GetTravelByID(id string) (*models.TravelInfo, error)
	GetTravelByOwner(owner int) ([]models.TravelInfo, error)
	GetTravelByLocation(location string) ([]models.TravelInfo, error)
	GetTravelByState(state int) ([]models.TravelInfo, error)
}

// TravelService is interface for travel service
type TravelService interface {
	travelGetter
	travelUpdater

	CreateTravel(*models.TravelInfo) error
	DeleteTravel(id string) error
}

type travelId = string
type travelMap = map[travelId]*models.TravelInfo

type travelServiceImpl struct {
	repo      repository.TravelRepository
	tempCache travelMap
}

func (t *travelServiceImpl) GetTravelByID(id string) (*models.TravelInfo, error) {

}

func (t *travelServiceImpl) GetTravelByOwner(owner int) ([]models.TravelInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) GetTravelByLocation(location string) ([]models.TravelInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) GetTravelByState(state int) ([]models.TravelInfo, error) {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateState(id string, state int) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateLocation(id string, location string) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateDetails(id string, details string) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateTogetherWith(id string, togetherWith []int) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateStartTime(id string, startTime int64) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateEndTime(id string, endTime int64) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateOwner(id string, owner int) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) CreateTravel(info *models.TravelInfo) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) DeleteTravel(id string) error {
	//TODO implement me
	panic("implement me")
}

func NewTravelService(repo repository.TravelRepository) TravelService {
	return &travelServiceImpl{repo}
}
