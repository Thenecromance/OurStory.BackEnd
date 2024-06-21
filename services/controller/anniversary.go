package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/gin-gonic/gin"
)

type anniversaryRoutes struct {
	base Interface.IRoute
	list Interface.IRoute
}

type AnniversaryController struct {
	anniversaryRoutes
	services services.AnniversaryService
}

func (c *AnniversaryController) Initialize() {
	c.base = route.NewREST("/api/anniversary")
	{
		c.base.SetHandler(
			c.getAnniversary,
			c.createAnniversary,
			c.updateAnniversary,
			c.deleteAnniversary,
		)
	}

	c.list = route.NewREST("/api/anniversary/list")
	{
		c.list.SetHandler(
			c.getAnniversaries,
		)
	}

}

//---------------------------------------------------------
//handlers
//---------------------------------------------------------

func (c *AnniversaryController) getAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	/*	obj, err := c.services.GetAnniversaryById()
		if err != nil {
			log.Infof("failed to get anniversary by id with error: %s", err.Error())
			resp.Error("failed to get anniversary by id with error")
			return
		}

		resp.Success(obj)*/

	panic("implement me")
}

func (c *AnniversaryController) createAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}
func (c *AnniversaryController) updateAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}
func (c *AnniversaryController) deleteAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

// ---------------------------------------------------------
// anniversary List handlers
// ---------------------------------------------------------
func (c *AnniversaryController) getAnniversaries(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func (c *AnniversaryController) getAnniversaryById(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewAnniversaryController(services services.AnniversaryService) *AnniversaryController {
	c := new(AnniversaryController)
	c.services = services
	c.Initialize()
	return c
}
