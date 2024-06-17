package repository

type TravelRepository interface {
	CreateTravel() error
	DeleteTravel() error
	UpdateTravel() error
	GetTravel() error
}

type travelRepository struct {
}

//
//func (tr *travelRepository) GetTravel() *TravelRepository {
//	return tr
//}
