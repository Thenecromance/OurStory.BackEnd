package repository

import (
	"fmt"
	"github.com/Thenecromance/OurStories/application/models"
	"gopkg.in/gorp.v2"
)

/*
	type DevShopRepository interface {
		// this method will return all items in data base for development purposes
		DevGetAllItems() []models.Item
	}
*/
type ShopRepository interface {
	// method will return all available items
	GetAllItems() []models.Item

	AddItem(item models.Item) (int, error)

	UpdateItem(item models.Item) error

	DeleteItem(id int) error
}

type CartRepository interface {
	// if user does not have a cart, a new cart will be created and the cart id will be returned
	HasCart(uid int64) bool

	GetCart(uid int64) (int64, error)

	CreateCart(uid int64) (int64, error)

	AddItemIntoCart(cartId int64, itemId int64, count int) error

	CleanCart(cartId int64) error
}

func NewShopRepository(db *gorp.DbMap) ShopRepository {
	obj := &shopRepository{
		db: db,
	}
	obj.initTable()
	return obj
}

type shopRepository struct {
	db *gorp.DbMap
}

func (s *shopRepository) initTable() error {
	s.db.AddTableWithName(models.Item{}, "item").SetKeys(true, "ID")

	return s.db.CreateTablesIfNotExists()
}

func (s *shopRepository) GetAllItems() []models.Item {
	var items []models.Item
	_, err := s.db.Select(&items, "SELECT * FROM item")
	if err != nil {
		return nil
	}

	return items
}

func (s *shopRepository) AddItem(item models.Item) (int, error) {
	err := s.db.Insert(&item)
	if err != nil {
		return 0, err
	}
	return item.ID, nil
}

func (s *shopRepository) UpdateItem(item models.Item) error {
	update, err := s.db.Update(&item)
	if err != nil {
		return err
	}
	if update == 0 {
		return fmt.Errorf("no item with id %d", item.ID)
	}
	return nil
}

func (s *shopRepository) DeleteItem(id int) error {
	_, err := s.db.Exec("DELETE FROM item WHERE id = ?", id)
	return err
}

//======================================================================================================================
// CartRepository
//======================================================================================================================

func NewCartRepository(db *gorp.DbMap) CartRepository {
	obj := &cartRepository{
		db: db,
	}
	obj.initTable()
	return obj
}

type cartRepository struct {
	db *gorp.DbMap
}

func (c *cartRepository) initTable() error {
	c.db.AddTableWithName(models.Cart{}, "cart").SetKeys(true, "ID")
	return c.db.CreateTablesIfNotExists()
}

func (c *cartRepository) HasCart(uid int64) bool {

	selectInt, err := c.db.SelectInt("SELECT COUNT(*) FROM cart WHERE user_id = ?", uid)
	if err != nil {
		return false
	}
	return selectInt > 0
}

func (c *cartRepository) GetCart(uid int64) (int64, error) {
	var cart models.Cart
	_, err := c.db.Select(&cart, "SELECT * FROM cart WHERE user_id = ?", uid)
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func (c *cartRepository) CreateCart(uid int64) (int64, error) {
	cart := models.Cart{
		UserId: uid,
	}
	err := c.db.Insert(&cart)
	if err != nil {
		return 0, err
	}
	return cart.ID, nil
}

func (c *cartRepository) AddItemIntoCart(cartId int64, itemId int64, count int) error {
	obj, err := c.db.Get(models.Item{}, cartId)
	if err != nil {
		return err
	}
	cart := obj.(*models.Cart)
	for i, item := range cart.Items {
		if item.Id == itemId {
			cart.Items[i].Count += count
		}
	}
	_, err = c.db.Update(cart)
	return err
}

func (c *cartRepository) CleanCart(cartId int64) error {

	i, err := c.db.Delete(models.Cart{ID: cartId})
	if err != nil {
		return err
	}
	if i == 0 {
		return fmt.Errorf("no cart with id %d", cartId)
	}
	return nil
}
