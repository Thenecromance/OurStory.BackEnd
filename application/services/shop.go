package services

import (
	"github.com/Thenecromance/OurStories/application/repository"
	"github.com/Thenecromance/OurStories/utility/log"
)

type CartService interface {
	// if user does not have a cart, a new cart will be created and the cart id will be returned
	GetCart(uid int64) (int64, error)
	CreateCart(uid int64) (int64, error)
	AddItemIntoCart(cartId int64, itemId int64, count int) error
	RemoveItemFromCart(cartId int64, itemId int64, count int) error
	// in this method, cart will be cleaned and all items will be removed from the cart
	// also this cart id will be deprecated
	CleanCart(cartId int64) error
}

// CheckoutService is an interface for checking out the cart
type CheckoutService interface {
	Checkout(cartId int64) error

	Result(cartId int64) (string, error)
}

type ShopService interface {
	// cart service
	CartService

	CheckoutService
}

func NewShopService(shop_ repository.ShopRepository, cart_ repository.CartRepository) ShopService {
	return &shopService{
		shop: shop_,
		cart: cart_,
	}
}

type shopService struct {
	shop repository.ShopRepository
	cart repository.CartRepository

	// cart id -> item id map -> item count
	//itemCache map[int64] /*cart id*/ map[int64]int
}

func (s *shopService) GetCart(uid int64) (int64, error) {
	if s.cart.HasCart(uid) {
		return s.cart.GetCart(uid)
	} else {

		log.Debugf("User %d does not have a cart, creating a new cart", uid)
		return s.cart.CreateCart(uid)
	}
}

func (s *shopService) CreateCart(uid int64) (int64, error) {
	return s.cart.GetCart(uid) // temp just let user could only has one cart at a time
}

func (s *shopService) AddItemIntoCart(cartId int64, itemId int64, count int) error {
	/*	if _, exists := s.itemCache[cartId]; !exists {
			s.itemCache[cartId] = make(map[int64]int)
		}

		s.itemCache[cartId][itemId] += count
	*/

	return s.cart.AddItemIntoCart(cartId, itemId, count)
}

func (s *shopService) RemoveItemFromCart(cartId int64, itemId int64, count int) error {
	return s.cart.AddItemIntoCart(cartId, itemId, -count)
}

func (s *shopService) CleanCart(cartId int64) error {

	return s.cart.CleanCart(cartId)
}

func (s *shopService) Checkout(cartId int64) error {
	//TODO implement me
	panic("implement me")
}

func (s *shopService) Result(cartId int64) (string, error) {
	//TODO implement me
	panic("implement me")
}
