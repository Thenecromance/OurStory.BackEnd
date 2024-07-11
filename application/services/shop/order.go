package shop

// Processor is an interface that defines the methods that a order processor should implement
// if implement this interface, the processor should be
type OrderProcessor interface {
	CreateOrder(uid int64) (int64, error)

	CancelOrder(orderID int64) error
}
