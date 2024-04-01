package User

import "errors"

type Model struct {
}

func (m *Model) init() {
	BindInfoTable()
}
func (m *Model) authorize(i *Info) bool {
	authUser := i.Copy()

	err := authUser.GetFromSQLByUserName()
	if err != nil {
		return false
	}

	if authUser.UserName != i.UserName ||
		authUser.Password != i.Password {
		return false
	}

	return true
}

func (m *Model) authUser(i *Info) error {

	//if i.Id == 0 {
	//	return errors.New("User not exist")
	//}
	if i.UserName == "" {
		return errors.New("User name is empty")
	}
	if i.Password == "" {
		return errors.New("Password is empty")
	}
	if !m.authorize(i) {
		return errors.New("User name or password is wrong")
	}

	return nil
}

func (m *Model) register(i *Info) error {
	if i.UserName == "" {
		return errors.New("User name is empty")
	}
	if i.Password == "" {
		return errors.New("Password is empty")
	}

	i.InsertToSQL()
	return nil
}

func (m *Model) updateProfile(i *Info) error {

	return nil
}
