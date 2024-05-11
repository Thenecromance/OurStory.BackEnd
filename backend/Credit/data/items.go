package data

// Item is the struct of the item which user can buy in the store
type Item struct {
	Id     int    `json:"id"`     // Item ID
	Name   string `json:"name"`   // Item name
	Price  int    `json:"price"`  // The price of the item
	Enable bool   `json:"enable"` // Whether the item is enabled
	Count  int    `json:"count"`  // The number of items in stock
}
