package constants

const (
	UnPurchase = iota // default state when the order is created
	Purchased         // all purchase sequence is done, this order will be marked as this state
	Cancelled         // user cancel the order, this order will be marked as this state
	Refunded          // the order is refunded, this order will be marked as this state
)
