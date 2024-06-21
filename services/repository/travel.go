package repository

import "github.com/Thenecromance/OurStories/services/models"

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
}

//
//func (tr *travelRepository) GetTravel() *TravelRepository {
//	return tr
//}
