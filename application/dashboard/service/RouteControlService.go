package service

import "github.com/Thenecromance/OurStories/application/dashboard/models"

type RouteControlService interface {
	GetRoutesByRole(role int) ([]models.RouteDTO, error)
}
