package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type AnniversaryService struct {
	repo repository.Anniversary
}

func (as *AnniversaryService) CreateAnniversary(anniversary *models.Anniversary) error {
	return as.repo.CreateAnniversary(anniversary)
}

func (as *AnniversaryService) RemoveAnniversary(anniversary *models.Anniversary) error {
	return as.repo.RemoveAnniversary(anniversary)
}

func (as *AnniversaryService) UpdateAnniversary(anniversary *models.Anniversary) error {
	return as.repo.UpdateAnniversary(anniversary)
}

func (as *AnniversaryService) GetAnniversaryById(id int64) (*models.Anniversary, error) {
	return as.repo.GetAnniversaryById(id)
}

func (as *AnniversaryService) GetAnniversaryList(user string) ([]models.Anniversary, error) {
	return as.repo.GetAnniversaryList(user)
}

func NewAnniversaryService(repo repository.Anniversary) *AnniversaryService {
	return &AnniversaryService{repo}
}
