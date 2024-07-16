package services

import (
	"fmt"
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/utility/log"
	"strconv"
)

type AnniversaryService interface {
	GetAnniversaryById(userid, id int) (*models.Anniversary, error)
	CreateAnniversary(userid int64, anniversary *models.Anniversary) error
	RemoveAnniversary(userid int, id int) error
	UpdateAnniversary(userid int, anniversary *models.Anniversary) error

	GetAnniversaries(userId int64) ([]models.Anniversary, error)
}

type anniversaryServiceImpl struct {
	repo repository.Anniversary
}

func (a *anniversaryServiceImpl) GetAnniversaryById(userid, id int) (*models.Anniversary, error) {

	anniversary, err := a.repo.GetAnniversaryById(userid, id)
	if err != nil {
		log.Debugf("error getting anniversary by id")
		return nil, err
	}

	if anniversary == nil {
		return nil, nil
	}

	return anniversary, nil
}

func (a *anniversaryServiceImpl) CreateAnniversary(userid int64, anniversary *models.Anniversary) error {
	if userid != anniversary.UserId {
		return fmt.Errorf("user id not match %d != %d", userid, anniversary.UserId)
	}

	err := a.repo.CreateAnniversary(anniversary)
	return err
}

func (a *anniversaryServiceImpl) RemoveAnniversary(userid int, id int) error {

	//a.repo.RemoveAnniversary()

	return a.repo.RemoveAnniversaryById(userid, id)

}

func (a *anniversaryServiceImpl) UpdateAnniversary(userid int, anniversary *models.Anniversary) error {
	return a.repo.UpdateAnniversary(anniversary)
}

func (a *anniversaryServiceImpl) GetAnniversaries(userId int64) ([]models.Anniversary, error) {

	// format int64 to string
	user := strconv.FormatInt(userId, 10)

	anniversaries, err := a.repo.GetAnniversaryList(user)
	if err != nil {
		log.Error("error getting anniversary list")
		return nil, err
	}
	var result []models.Anniversary
	for _, v := range anniversaries {
		anni := v
		result = append(result, anni)
	}
	return result, nil
}

func NewAnniversaryService(repo repository.Anniversary) AnniversaryService {
	return &anniversaryServiceImpl{repo}
}
