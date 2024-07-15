package services

import (
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/application/services/pwdHashing"
	"github.com/Thenecromance/OurStories/application/services/pwdHashing/Scrypt"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	"github.com/Thenecromance/OurStories/utility/log"
	"time"
)

type UserService interface {
	GetUserIdByName(username string) (int64, error)
	GetUser(id int) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	AuthorizeUser(login *models.UserLogin) (bool, error)
	SignedTokenToUser(info interface{}) string
	AddUser(user *models.UserRegister) error
	HasUserAndEmail(username, email string) bool
}

type userServiceImpl struct {
	repo repository.UserRepository
	auth pwdHashing.PwdHasher
}

func (us *userServiceImpl) GetUserIdByName(username string) (int64, error) {
	return us.repo.GetUserIdByName(username)
}

func (us *userServiceImpl) GetUser(id int) (*models.User, error) {
	return us.repo.GetUser(id)
}

func (us *userServiceImpl) GetUserByUsername(username string) (*models.User, error) {
	return us.repo.GetUserByUsername(username)
}

func (us *userServiceImpl) AuthorizeUser(login *models.UserLogin) (bool, error) {

	usrInDb, err := us.repo.GetUserByUsername(login.UserName)
	if err != nil {
		log.Error("error in getting user", err)
		return false, err
	}
	if usrInDb == nil {
		log.Error("user not found")
		return false, nil
	}

	if us.auth.Verify(login.Password, usrInDb.Password, usrInDb.Salt) {
		err = us.repo.UpdateLastLogin(usrInDb.UserId, time.Now().Unix())
		if err != nil {
			log.Error("error in updating last login", err)
			return false, err
		}
		return true, nil
	}

	/*if hash.Compare(login.Password, usrInDb.Password) {

	}*/
	return false, nil
}

func (us *userServiceImpl) SignedTokenToUser(info interface{}) string {
	token, err := JWT.Instance().SignTokenToUser(info)
	if err != nil {
		log.Error("error in signing token", err)
		return ""
	}
	return token
}

func (us *userServiceImpl) AddUser(user *models.UserRegister) error {

	hashedPwd, hashedSalt := us.auth.Hash(user.Password)
	fullInfo := &models.User{
		UserAdvancedDTO: models.UserAdvancedDTO{
			UserBasicDTO: models.UserBasicDTO{
				UserName: user.UserName,
			},
			Email: user.Email,
		},
		Password: hashedPwd,
		Salt:     hashedSalt,
	}

	fullInfo.CreatedAt = time.Now().Unix()
	fullInfo.LastLogin = fullInfo.CreatedAt

	return us.repo.InsertUser(fullInfo)
}

func (us *userServiceImpl) HasUserAndEmail(username, email string) bool {
	return us.repo.HasUserAndEmail(username, email)
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo, Scrypt.New()}
}
