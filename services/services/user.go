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
func (us *UserService) AuthorizeUser(login *models.UserLogin) (models.User, error) {
	panic("token has not been implemented")
	return models.User{}, nil
}

func (us *UserService) SignedTokenToUser(info string) string {
	panic("token has not been implemented")
	return ""
}

func (us *UserService) AddUser(user *models.UserRegister) error {
	fullInfo := &models.User{
		UserAdvancedDTO: models.UserAdvancedDTO{
			UserBasicDTO: models.UserBasicDTO{
				UserName: user.UserName,
			},
			Email: user.Email,
		},
		Password: user.Password,
	}
	return us.repo.InsertUser(fullInfo)
}

func (us *UserService) HasUserAndEmail(username, email string) bool {
	return us.repo.HasUserAndEmail(username, email)
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}
