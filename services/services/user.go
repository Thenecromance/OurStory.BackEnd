package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (us *UserService) GetUser(id int) (models.User, error) {
	return us.repo.GetUser(id)
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}
