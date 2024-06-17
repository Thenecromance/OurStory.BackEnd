package services

import "github.com/Thenecromance/OurStories/application/repository"

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}
