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

func NewTravelService(repo repository.TravelRepository) TravelService {
	return &travelServiceImpl{repo}
}
