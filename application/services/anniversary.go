package services

import (
	"github.com/Thenecromance/OurStories/application/model"
	"github.com/Thenecromance/OurStories/application/repository"
)

type AnniversaryService struct {
	repo repository.Anniversary
}

func (as *AnniversaryService) CreateAnniversary(anniversary *model.Anniversary) error {
	return as.repo.CreateAnniversary(anniversary)
}

func (as *AnniversaryService) RemoveAnniversary(anniversary *model.Anniversary) error {
	return as.repo.RemoveAnniversary(anniversary)
}

func (as *AnniversaryService) UpdateAnniversary(anniversary *model.Anniversary) error {
	return as.repo.UpdateAnniversary(anniversary)
}

func (as *AnniversaryService) GetAnniversaryById(id int64) (*model.Anniversary, error) {
	return as.repo.GetAnniversaryById(id)
}

func (as *AnniversaryService) GetAnniversaryList(user string) ([]model.Anniversary, error) {
	return as.repo.GetAnniversaryList(user)
}

func NewAnniversaryService(repo repository.Anniversary) *AnniversaryService {
	return &AnniversaryService{repo}
}
