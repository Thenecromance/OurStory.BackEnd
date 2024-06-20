package services

import "github.com/Thenecromance/OurStories/services/repository"

type TravelService interface {
	GetTravel(id string) (any, error)
	CreateTravel(any) error
	UpdateTravel(any) error
	DeleteTravel(id string) error
}

type travelServiceImpl struct {
	repo repository.TravelRepository
}

func (t *travelServiceImpl) GetTravel(id string) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) CreateTravel(a any) error {
	//TODO implement me
	panic("implement me")
}

func (t *travelServiceImpl) UpdateTravel(a any) error {
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
