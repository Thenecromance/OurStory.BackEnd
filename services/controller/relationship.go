package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/services/models"
	"github.com/Thenecromance/OurStories/services/services"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

// the relationship controller should have the following routes
// 1. createLink - POST /api/relation
// 2. activeLink - GET /api/relation/:id - get the user's associate link just like: https://m0nkeycl1cker.com/api/relation/d8290f28049f4f538d4df2a10a922a3783c1ee3df0e15e9632058f0d06a0639c
// 3. deleteLink - DELETE /api/relation/:id - delete the user's associate link
// 4. getRelation - GET /api/relation/:user - get the user associate with the other's list
// 5. associateHistory - GET /api/relation/:user/history - get the user's associate history

type relationGroups struct {
	createLink       Interface.IRoute
	activeLink       Interface.IRoute
	deleteLink       Interface.IRoute
	getRelation      Interface.IRoute
	associateHistory Interface.IRoute
}

type relationshipController struct {
	relationGroups
	service services.RelationShipService
}

func (r *relationshipController) Initialize() {
	r.SetRoutes()
}

func (r *relationshipController) Name() string {
	return "relationship"
}

func (r *relationshipController) SetRoutes() {
	r.createLink = route.NewRouter("/api/relation", "POST")
	{
		r.createLink.SetHandler(r.createBindLink)
	}
	r.activeLink = route.NewRouter("/api/relation/:id", "GET")
	{
		r.activeLink.SetHandler(r.linkUser)
	}
	r.deleteLink = route.NewRouter("/api/relation/:id", "DELETE")
	{
		r.deleteLink.SetHandler(r.unbindUser)
	}
	r.getRelation = route.NewRouter("/api/relation/:user", "GET")
	{
	}
	r.associateHistory = route.NewRouter("/api/relation/:user/history", "GET")
	{
	}
}

func (r *relationshipController) GetRoutes() []Interface.IRoute {
	return []Interface.IRoute{r.createLink, r.activeLink, r.deleteLink, r.getRelation, r.associateHistory}
}

// ---------------------------------------------------------
// handlers
// ---------------------------------------------------------

// @Title createBindLink the user's associate link
// @description get the user's associate link
func (r *relationshipController) createBindLink(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type request struct {
		UserID       int `json:"user_id,omitempty" form:"user_id"`
		RelationType int `json:"relation_type,omitempty" form:"relation_type"`
	}

	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error(err)
		resp.Error("invalid request")
		return
	}

	// create the relationship
	link := r.service.CreateRelationshipConnection(req.UserID, req.RelationType)
	if link == "" {
		resp.Error("failed to create the relationship, you have reached the limit")
		return
	}

	var urlResp models.RelationShipResponse
	urlResp.URL = link
	urlResp.RelationType = req.RelationType
	resp.Success(urlResp)
}

// @Title linkUser
// @description start to process the user's associate link jobs
func (r *relationshipController) linkUser(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	// get the user's associate link
	link := ctx.Param("id")
	if link == "" {
		resp.Error("invalid link")
		return
	}

}

func (r *relationshipController) unbindUser(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewRelationshipController(service services.RelationShipService) Interface.IController {
	return &relationshipController{
		service: service,
	}
}
