package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type AnniversaryService interface {
	GetAnniversaryById(id int64) (*models.Anniversary, error)
	CreateAnniversary(anniversary *models.Anniversary) error
	RemoveAnniversary(anniversary *models.Anniversary) error
	UpdateAnniversary(anniversary *models.Anniversary) error

	GetAnniversaries(user string) ([]models.Anniversary, error)
}

type anniversaryServiceImpl struct {
	repo repository.Anniversary
}

/*func NewAnniversaryService(repo repository.Anniversary) AnniversaryService {
	return &anniversaryServiceImpl{repo}
}
*/
