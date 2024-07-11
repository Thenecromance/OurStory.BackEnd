package services

type ShopService interface {
	// if user does not have a cart, a new cart will be created and the cart id will be returned
	GetCart(uid int64) (int64, error)
	CreateCart(uid int64) (int64, error)
	AddItemIntoCart(cartId int64, itemId int64, count int) error
	RemoveItemFromCart(cartId int64, itemId int64, count int) error
	// in this method, cart will be cleaned and all items will be removed from the cart
	// also this cart id will be deprecated
	CleanCart(cartId int64) error

	// Checkout is a method that should be called when the user wants to checkout the cart
	Checkout(cartId int64) error
}

func NewShopService() ShopService {
	return &shopService{}
}

type shopService struct {
}
