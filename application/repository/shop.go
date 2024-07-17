package repository

import (
	"fmt"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/utility/log"
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

	AddItem(item models.Item) (int64, error)

	UpdateItem(item models.Item) error

	DeleteItem(id int64) error

	Interface.Repository
}

type CartRepository interface {
	Interface.Repository

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
	//obj.initTable()
	return obj
}

type shopRepository struct {
	db *gorp.DbMap
}

func (s *shopRepository) BindTable() error {
	s.db.AddTableWithName(models.Item{}, "Items")
	s.db.AddTableWithName(models.Transaction{}, "Transactions")
	s.db.AddTableWithName(models.TransactionLog{}, "TransactionLogs")
	s.db.AddTableWithName(models.UserBalance{}, "UserBalances").SetKeys(false, "UserId")
	return nil
}

/*
func (s *shopRepository) initTable() error {
	s.db.AddTableWithName(models.Item{}, "item").SetKeys(true, "CartId")

	return s.db.CreateTablesIfNotExists()
}*/

func (s *shopRepository) GetAllItems() []models.Item {
	var items []models.Item
	_, err := s.db.Select(&items, "SELECT * FROM item")
	if err != nil {
		return nil
	}

	return items
}

func (s *shopRepository) AddItem(item models.Item) (int64, error) {
	err := s.db.Insert(&item)
	if err != nil {
		return 0, err
	}
	return item.ItemId, nil
}

func (s *shopRepository) UpdateItem(item models.Item) error {
	update, err := s.db.Update(&item)
	if err != nil {
		return err
	}
	if update == 0 {
		return fmt.Errorf("no item with id %d", item.ItemId)
	}
	return nil
}

func (s *shopRepository) DeleteItem(id int64) error {
	i, err := s.db.Delete(models.Item{ItemId: id})
	log.Debugf("delete item %d", i)
	return err
}

//======================================================================================================================
// CartRepository
//======================================================================================================================

func NewCartRepository(db *gorp.DbMap) CartRepository {
	obj := &cartRepository{
		db: db,
	}
	//obj.initTable()
	return obj
}

type cartRepository struct {
	db *gorp.DbMap
}

func (c *cartRepository) BindTable() error {
	c.db.AddTableWithName(models.Carts{}, "Carts")
	c.db.AddTableWithName(models.CartedItem{}, "CartedItems")
	return nil
}

/*func (c *cartRepository) initTable() error {
	c.db.AddTableWithName(models.Carts{}, "cart").SetKeys(true, "CartId")
	return c.db.CreateTablesIfNotExists()
}*/

func (c *cartRepository) HasCart(uid int64) bool {

	selectInt, err := c.db.SelectInt("SELECT COUNT(*) FROM Carts WHERE user_id = ?", uid)
	if err != nil {
		return false
	}
	return selectInt > 0
}

func (c *cartRepository) GetCart(uid int64) (int64, error) {

	var cart models.Carts
	err := c.db.SelectOne(&cart, "SELECT * FROM Carts WHERE user_id = ?", uid)
	if err != nil {
		return 0, err
	}

	return cart.CartId, nil
}

func (c *cartRepository) CreateCart(uid int64) (int64, error) {
	cart := models.Carts{
		UserId: uid,
	}
	err := c.db.Insert(&cart)
	if err != nil {
		return 0, err
	}
	return cart.CartId, nil
}

func (c *cartRepository) AddItemIntoCart(cartId int64, itemId int64, count int) error {
	obj, err := c.db.Get(models.Item{}, cartId)
	if err != nil {
		return err
	}
	cart := obj.(*models.Carts)
	/*for i, item := range cart.Items {
		if item.UserId == itemId {
			cart.Items[i].Count += count
		}
	}*/
	_, err = c.db.Update(cart)
	return err
}

func (c *cartRepository) CleanCart(cartId int64) error {

	i, err := c.db.Delete(models.Carts{CartId: cartId})
	if err != nil {
		return err
	}
	if i == 0 {
		return fmt.Errorf("no cart with id %d", cartId)
	}
	return nil
}
