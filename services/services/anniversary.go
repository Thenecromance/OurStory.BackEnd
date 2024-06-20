package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type AnniversaryService interface {
	CreateAnniversary(anniversary *models.Anniversary) error
	GetAnniversaryById(id int64) (*models.Anniversary, error)
	UpdateAnniversary(anniversary *models.Anniversary) error
	RemoveAnniversary(anniversary *models.Anniversary) error
}

type anniversaryServiceImpl struct {
	repo repository.Anniversary
}

func (as *anniversaryServiceImpl) CreateAnniversary(anniversary *models.Anniversary) error {
	return as.repo.CreateAnniversary(anniversary)
}

func (as *anniversaryServiceImpl) RemoveAnniversary(anniversary *models.Anniversary) error {
	return as.repo.RemoveAnniversary(anniversary)
}

func (as *anniversaryServiceImpl) UpdateAnniversary(anniversary *models.Anniversary) error {
	return as.repo.UpdateAnniversary(anniversary)
}

func (as *anniversaryServiceImpl) GetAnniversaryById(id int64) (*models.Anniversary, error) {
	return as.repo.GetAnniversaryById(id)
}

func (as *anniversaryServiceImpl) GetAnniversaryList(user string) ([]models.Anniversary, error) {
	return as.repo.GetAnniversaryList(user)
}

func NewAnniversaryService(repo repository.Anniversary) AnniversaryService {
	return &anniversaryServiceImpl{repo}
}
