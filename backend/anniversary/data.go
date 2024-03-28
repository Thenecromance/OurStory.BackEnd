package anniversary

type date struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Date        int    `json:"date" db:"date"`
}
