package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/services/services"
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
	//TODO implement me
	panic("implement me")
}

func (r *relationshipController) Name() string {
	return "relationship"
}

func (r *relationshipController) SetRoutes() {
	r.createLink = route.NewRouter("/api/relation", "POST")
	{
		r.createLink.SetHandler(r.getAssociateLink)
	}
	r.activeLink = route.NewRouter("/api/relation/:id", "GET")
	{
	}
	r.deleteLink = route.NewRouter("/api/relation/:id", "DELETE")
	{
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

// @Title getAssociateLink the user's associate link
// @description get the user's associate link
func (r *relationshipController) getAssociateLink(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	/*	if r.service.UserHasAssociation() {

		}*/
}

func (r *relationshipController) associateWithOther(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)
}

func NewRelationshipController(service services.RelationShipService) Interface.IController {
	return &relationshipController{
		service: service,
	}
}
