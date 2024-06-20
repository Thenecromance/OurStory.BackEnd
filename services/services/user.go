package services

import (
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/repository"
	"github.com/Thenecromance/OurStories/utility/hash"
	"github.com/Thenecromance/OurStories/utility/log"
)

type UserService interface {
	GetUserIdByName(username string) (int, error)
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	AuthorizeUser(login *models.UserLogin) (*models.User, error)
	SignedTokenToUser(info interface{}) string
	AddUser(user *models.UserRegister) error
	HasUserAndEmail(username, email string) bool
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func (us *userServiceImpl) GetUserIdByName(username string) (int, error) {
	return us.repo.GetUserIdByName(username)
}

func (us *userServiceImpl) GetUser(id int) (*models.User, error) {
	return us.repo.GetUser(id)
}

func (us *userServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	return us.repo.GetUserByUsername(username)
}
func (us *userServiceImpl) AuthorizeUser(login *models.UserLogin) (*models.User, error) {
	return us.repo.GetUserByUsername(login.UserName)
}

func (us *userServiceImpl) SignedTokenToUser(info interface{}) string {
	token, err := JWT.Instance().AuthorizeUser(info)
	if err != nil {
		log.Error("error in signing token", err)
		return ""
	}
	return token
}

func (us *userServiceImpl) AddUser(user *models.UserRegister) error {
	fullInfo := &models.User{
		UserAdvancedDTO: models.UserAdvancedDTO{
			UserBasicDTO: models.UserBasicDTO{
				UserName: user.UserName,
			},
			Email: user.Email,
		},
		Password: hash.Hash(user.Password),
	}
	return us.repo.InsertUser(fullInfo)
}

func (us *userServiceImpl) HasUserAndEmail(username, email string) bool {
	return us.repo.HasUserAndEmail(username, email)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo}
}
