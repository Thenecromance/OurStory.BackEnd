package User

import (
	"fmt"
	"github.com/Thenecromance/OurStories/backend/User/Data"
	"github.com/Thenecromance/OurStories/backend/User/Validator"
	"github.com/Thenecromance/OurStories/backend/User/creator"
	"github.com/Thenecromance/OurStories/base/log"
	"github.com/Thenecromance/OurStories/base/lru"
	"time"
)

type Model struct {
	cache     *lru.Cache
	creator   *creator.Creator
	validator *Validator.Validator
}

// initialize the model
func (m *Model) init() {
	Data.BindInfoTable()

	m.cache = lru.New(100)
	m.creator = creator.New()
	m.validator = Validator.New()

}

func (m *Model) RequestFromDatabase(username string) (*Data.Info, error) {
	usr := &Data.Info{
		AuthorizationInfo: Data.AuthorizationInfo{
			UserName: username,
		},
	}
	err := usr.GetFromSQLByUserName()
	if err != nil {
		return nil, fmt.Errorf("could not find user: %s", username)
	}
	return usr, nil
}

func (m *Model) Login(data *Data.AuthorizationInfo) (*Data.Info, error) {
	var err error
	info, ok := m.cache.Get(data.UserName)
	if !ok {
		info, err = m.RequestFromDatabase(data.UserName)
	}

	if info == nil || err != nil {
		return nil, fmt.Errorf("could not find user: %s", data.UserName)
	}

	if !m.validator.ValidateUser(data, info.(*Data.Info)) {
		return nil, fmt.Errorf("username or password is wrong")
	}

	//valid user, update last login time

	//update(or add to cache)
	m.cache.Add(info.(*Data.Info).UserName, info, time.Now().Add(time.Hour*24))

	return info.(*Data.Info), nil
}

func (m *Model) Register(data *Data.AuthorizationInfo, email string) error {
	_, ok := m.cache.Get(data.UserName)
	if ok {
		return fmt.Errorf("username already exists")
	}

	err := m.creator.NewUser(data.UserName, email, data.Password)
	if err != nil {
		return err
	}

	usr := &Data.Info{
		AuthorizationInfo: *data,
	}

	_ = usr.GetFromSQLByUserName()

	m.cache.Add(data.UserName, usr, time.Now().Add(time.Hour*24))

	return nil
}

func (m *Model) Profile(username string) (*Data.Info, error) {
	//check if the user is logged in
	info, ok := m.cache.Get(username)

	if !ok {
		return nil, fmt.Errorf("please login first")
	}
	return info.(*Data.Info), nil
}

func (m *Model) UpdateProfile(username string, newInfo *Data.CommonInfo) (*Data.Info, error) {
	info, err := m.Profile(username)

	if err != nil {
		log.Errorf("update profile failed: %s", err)
		return nil, err
	}
	info.ApplyNewInfo(newInfo)

	err = info.UpdateToSQL()
	if err != nil {
		log.Errorf("update profile failed: %s", err)
		return nil, fmt.Errorf("update profile failed: %s", err)
	}

	m.cache.Add(username, info, time.Now().Add(time.Hour*24))
	return info, nil
}

func (m *Model) LogoutUser(username string) {
	m.cache.Remove(username)
}
