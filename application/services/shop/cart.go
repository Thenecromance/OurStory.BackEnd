package shop

type CartService interface {
	GetCart(userId int64) (int64, error)

	AddItemToCart(cartId int64, itemId int64, count int) error
}
