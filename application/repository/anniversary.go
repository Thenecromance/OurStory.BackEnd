package repository

import (
	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
	"gopkg.in/gorp.v2"
)

type Anniversary interface {
	Interface.Repository
	CreateAnniversary(anniversary *models.Anniversary) error
	RemoveAnniversary(anniversary *models.Anniversary) error
	RemoveAnniversaryById(userId, id int) error
	UpdateAnniversary(anniversary *models.Anniversary) error
	GetAnniversaryById(userId, id int) (*models.Anniversary, error)
	GetAnniversaryList(user string) ([]models.Anniversary, error)
}

type anniversaryRepository struct {
	db *gorp.DbMap
}

func (a *anniversaryRepository) CreateAnniversary(anniversary *models.Anniversary) error {
	trans, err := a.db.Begin()
	if err != nil {
		return err
	}
	err = trans.Insert(anniversary)
	if err != nil {
		trans.Rollback()
		return err
	}

	return trans.Commit()
}

func (a *anniversaryRepository) RemoveAnniversary(anniversary *models.Anniversary) error {
	trans, err := a.db.Begin()
	if err != nil {
		return err
	}
	id, err := trans.Delete(anniversary)
	if err != nil {
		trans.Rollback()
		return err
	}
	if id == 0 {
		//todo: add log
	}
	return trans.Commit()

}

func (a *anniversaryRepository) RemoveAnniversaryById(userId, id int) error {
	trans, err := a.db.Begin()
	if err != nil {
		return err
	}

	_, err = trans.Query("delete from anniversary where id = ? and user_id = ?", id, userId)
	if err != nil {
		trans.Rollback()
		return err
	}

	return trans.Commit()
}

func (a *anniversaryRepository) UpdateAnniversary(anniversary *models.Anniversary) error {
	trans, err := a.db.Begin()
	if err != nil {
		return err
	}
	update, err := trans.Update(anniversary)
	if err != nil {
		trans.Rollback()
		return err
	}
	log.Warnf("update anniversary with id: %d", update)
	return trans.Commit()
}

func (a *anniversaryRepository) GetAnniversaryById(userId, id int) (*models.Anniversary, error) {

	result := &models.Anniversary{}
	err := a.db.SelectOne(result, "select * from anniversary where id = ? and user_id = ?", id, userId)
	if err != nil {
		return nil, err
	}
	return result, nil

	//TODO implement me
	panic("implement me")
}

func (a *anniversaryRepository) GetAnniversaryList(userId string) ([]models.Anniversary, error) {
	result, err := a.db.Select(&models.Anniversary{}, "select * from anniversary where user_id = ?", userId)
	if err != nil {
		return nil, err
	}

	var anniversaries []models.Anniversary
	for _, v := range result {
		anniversaries = append(anniversaries, v.(models.Anniversary))
	}
	//TODO implement me
	panic("implement me")
}

/*
func (a *anniversaryRepository) initTable() {
	if a.db == nil {
		log.Error("db is nil")
		return
	}

	log.Info("start to binding anniversary with table travel")
	tbl := a.db.AddTableWithName(models.Anniversary{}, "anniversary")
	tbl.SetKeys(false, "UserId") // using snowflake to generate the id
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("UserId").SetNotNull(true)
	tbl.ColMap("Date").SetNotNull(true)
	tbl.ColMap("Name").SetNotNull(true)
	tbl.ColMap("Description").SetNotNull(true)
	tbl.ColMap("CreatedAt").SetNotNull(true)
	err := a.db.CreateTablesIfNotExists()
	if err != nil {
		log.Errorf("failed to create table anniversary with error: %s", err.Error())
		return
	}
}*/

func (a *anniversaryRepository) BindTable() error {
	a.db.AddTableWithName(models.Anniversary{}, "Anniversaries")
	return nil
}

func NewAnniversaryRepository(db *gorp.DbMap) Anniversary {
	a := &anniversaryRepository{db}
	return a
}
