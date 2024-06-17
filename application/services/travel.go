package services

import "github.com/Thenecromance/OurStories/application/repository"

type TravelService struct {
	repo repository.TravelRepository
}

func NewTravelService(repo repository.TravelRepository) *TravelService {
	return &TravelService{repo}
}
