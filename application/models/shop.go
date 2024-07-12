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

type Item struct {
	ID          int     `json:"id,omitempty" db:"id"`
	Name        string  `json:"name,omitempty" db:"name"`
	Detail      string  `json:"detail,omitempty" db:"detail"`
	Limit       int     `json:"limit,omitempty" db:"limit"`
	Price       float64 `json:"price,omitempty" db:"price"`
	ReleaseDate int64   `json:"release_date,omitempty" db:"release_date"`
	ExpireDate  int64   `json:"expire_date,omitempty" db:"expire_date"`
	CreateAt    int64   `json:"create_at,omitempty" db:"create_at"`
	Publisher   int     `json:"publisher,omitempty" db:"publisher"`
}

type Cart struct {
	ID     int64 // Cart id
	UserId int64 // user id
	Items  []struct {
		Id    int64 // id of the item
		Count int   // how many items of this type
	}
}
