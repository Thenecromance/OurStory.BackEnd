package repository

import "github.com/Thenecromance/OurStories/application/models"

type DevShopRepository interface {
	// this method will return all items in data base for development purposes
	DevGetAllItems() []models.Item
}

type ShopRepository interface {
	// method will return all available items
	GetAllItems() []models.Item

	AddItem(item models.Item) (int, error)

	UpdateItem(item models.Item) error

	DeleteItem(id int) error
}
