package services

import (
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
)

type UserService interface {
	GetUser(id int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	AuthorizeUser(login *models.UserLogin) (models.User, error)
	SignedTokenToUser(info string) string
	AddUser(user *models.UserRegister) error
	HasUserAndEmail(username, email string) bool
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func (us *userServiceImpl) GetUser(id int) (models.User, error) {
	return us.repo.GetUser(id)
}

func (us *userServiceImpl) GetUserByUsername(username string) (models.User, error) {
	return us.repo.GetUserByUsername(username)
}
func (us *userServiceImpl) AuthorizeUser(login *models.UserLogin) (models.User, error) {
	panic("token has not been implemented")
	return models.User{}, nil
}

func (us *userServiceImpl) SignedTokenToUser(info string) string {
	panic("token has not been implemented")
	return ""
}

func (us *userServiceImpl) AddUser(user *models.UserRegister) error {
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

func (us *userServiceImpl) HasUserAndEmail(username, email string) bool {
	return us.repo.HasUserAndEmail(username, email)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo}
}
