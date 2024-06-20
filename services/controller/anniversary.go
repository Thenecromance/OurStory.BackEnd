package controller

import "github.com/Thenecromance/OurStories/services/services"

type anniversaryRouter struct {
}

type AnniversaryController struct {
	services *services.AnniversaryService
}
