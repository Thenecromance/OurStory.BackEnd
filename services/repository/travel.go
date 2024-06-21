package repository

import "github.com/Thenecromance/OurStories/services/models"

type TravelRepository interface {
	CreateTravel(info *models.TravelInfo) error
	DeleteTravel(travelId int) error
	UpdateTravel(info *models.TravelInfo) error
	GetTravelByID(travelId int) (*models.TravelInfo, error)
}

type travelRepository struct {
}

//
//func (tr *travelRepository) GetTravel() *TravelRepository {
//	return tr
//}
