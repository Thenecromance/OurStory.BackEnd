package models

type Route struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Path      string `db:"path"`
	Component string `db:"component"` // this components must be a full path and must be a valid component
	AllowRole string `db:"allow_role"`
}

type RouteDTO struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"` // this components must be a full path and must be a valid component
}
