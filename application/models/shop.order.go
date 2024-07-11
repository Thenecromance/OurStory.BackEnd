package models

// Order is a struct that represents an order in the system
type Order struct {
	OrderID    int64
	CreatedAt  int64
	FinishedAt int64
	Items      []struct {
		Id    int64 // id of the item
		Count int   // how many items of this type
	}
	State int
}
