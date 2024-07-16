package controller

import (
	"strconv"
	"strings"

	"github.com/Thenecromance/OurStories/application/models"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	Interface2 "github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/response"
	"github.com/Thenecromance/OurStories/server/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type anniversaryRoutes struct {
	base Interface2.IRoute
	list Interface2.IRoute
}

type AnniversaryController struct {
	anniversaryRoutes
	services services.AnniversaryService
}

func (c *AnniversaryController) Name() string {
	return "AnniversaryController"
}

func (c *AnniversaryController) SetupRoutes() {

	mw := JWT.Middleware()
	c.base = route.NewREST("/api/anniversary")
	{
		c.base.SetMiddleWare(mw)
		c.base.SetHandler(
			c.getAnniversary,
			c.createAnniversary,
			c.updateAnniversary,
			c.deleteAnniversary,
		)
	}

	c.list = route.NewREST("/api/anniversary/list")
	{
		c.list.SetMiddleWare(mw)
		c.list.SetHandler(
			c.getAnniversaries,
			nil, nil, nil,
		)
	}
}

func (c *AnniversaryController) GetRoutes() []Interface2.IRoute {
	return []Interface2.IRoute{c.base, c.list}
}

func (c *AnniversaryController) Initialize() {
}

//---------------------------------------------------------
//handlers
//---------------------------------------------------------

func (c *AnniversaryController) getAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	id := ctx.Query("id")

	// todo: fucking this shit code, only shit can write this shit code
	var ids []string
	if strings.Contains(id, ",") {
		ids = strings.Split(id, ",")
		if len(ids) == 0 {
			resp.SetCode(response.BadRequest).AddData("Invalid id")
			return
		}

	} else {
		ids = []string{id}
	}

	for _, iid := range ids {
		anniId, err := strconv.Atoi(iid)
		if err != nil {
			log.Error(err)
			continue
		}
		obj, err := c.services.GetAnniversaryById(int(claim.Id), anniId)
		if err != nil {
			log.Warnf("failed to get anniversary with id: %d", anniId)
			continue
		}
		resp.AddData(obj)
	}
	if resp.Meta.Count == 0 {
		resp.SetCode(response.BadRequest).AddData("Invalid id")
		return
	} else {
		resp.SetCode(response.OK)
	}

	panic("implement me")
}

func (c *AnniversaryController) createAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	anni := models.Anniversary{}
	if err := ctx.BindJSON(&anni); err != nil {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	if claim.Id != anni.UserId {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	err := c.services.CreateAnniversary(claim.Id, &anni)

	if err != nil {
		log.Error(err)
		resp.SetCode(response.InternalServerError).AddData("Failed to create anniversary")
		return
	}
	resp.Success(anni)
}

func (c *AnniversaryController) updateAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	anni := models.Anniversary{}
	if err := ctx.BindJSON(&anni); err != nil {
		resp.SetCode(response.BadRequest).AddData("Invalid request")
		return
	}

	err := c.services.UpdateAnniversary(int(claim.Id), &anni)
	if err != nil {
		log.Error(err)
		resp.Error("Failed to update anniversary")
		return
	}
	resp.Success(anni)
}
func (c *AnniversaryController) deleteAnniversary(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	id := ctx.Query("id")
	if id == "" {
		resp.SetCode(response.BadRequest).AddData("Invalid id")
		return
	}

	var ids []string
	if strings.Contains(id, ",") {

		ids = strings.Split(id, ",")
		if len(ids) == 0 {
			resp.SetCode(response.BadRequest).AddData("Invalid id")
			return
		}
	} else {
		ids = []string{id}
	}

	for _, iid := range ids {
		anniId, err := strconv.Atoi(iid)
		if err != nil {
			log.Error(err)
			continue
		}
		err = c.services.RemoveAnniversary(int(claim.Id), anniId)
		if err != nil {
			log.Warnf("failed to delete anniversary with id: %d", anniId)
			continue
		}
	}

	resp.Success("success")
}

// ---------------------------------------------------------
// anniversary List handlers
// ---------------------------------------------------------
func (c *AnniversaryController) getAnniversaries(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	anniversaries, err := c.services.GetAnniversaries(claim.Id)
	if err != nil {
		log.Warn("failed to get anniversaries", err)
		resp.SetCode(response.BadRequest).AddData("Failed to get anniversaries")
		return
	}

	resp.Success(anniversaries)
}

func (c *AnniversaryController) getAnniversaryById(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewAnniversaryController(services services.AnniversaryService) Interface2.IController {
	c := new(AnniversaryController)
	c.services = services
	return c
}
