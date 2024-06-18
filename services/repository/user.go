package repository

import "github.com/Thenecromance/OurStories/services/models"

type UserRepository interface {
	GetUser(id int) (models.UserInDb, error)
	GetUsers() ([]models.UserInDb, error)
}
