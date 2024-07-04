package models

type Route struct {
	ID        int    `db:"id" json:"id,omitempty"`
	Name      string `db:"name" json:"name,omitempty"`
	Path      string `db:"path" json:"path,omitempty"`
	Component string `db:"component" json:"component,omitempty"` // this components must be a full path and must be a valid component
	AllowRole int    `db:"allow_role" json:"allow_role,omitempty"`
}

type RouteDTO struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"` // this components must be a full path and must be a valid component
}
