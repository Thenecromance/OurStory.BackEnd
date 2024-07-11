package shop

type CheckoutService interface {
	Checkout(cartId int64) error
}
