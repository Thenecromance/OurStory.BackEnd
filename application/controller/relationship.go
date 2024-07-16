package controller

import (
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/middleware/Authorization/JWT"
	Interface2 "github.com/Thenecromance/OurStories/server/Interface"
	"github.com/Thenecromance/OurStories/server/response"
	"github.com/Thenecromance/OurStories/server/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

// the relationship controller should have the following routes
// 1. createLink - POST /api/relation
// 2. activeLink - GET /api/relation/:id - get the user's associate link just like: https://m0nkeycl1cker.com/api/relation/d8290f28049f4f538d4df2a10a922a3783c1ee3df0e15e9632058f0d06a0639c
// 3. deleteLink - DELETE /api/relation/:id - delete the user's associate link
// 4. getRelation - GET /api/relation/list/:user - get the user associate with the other's list
// 5. associateHistory - GET /api/relation/history/:user - get the user's associate history

type relationGroups struct {
	createLink       Interface2.IRoute
	activeLink       Interface2.IRoute
	deleteLink       Interface2.IRoute
	getRelation      Interface2.IRoute
	associateHistory Interface2.IRoute
}

type relationshipController struct {
	relationGroups
	service services.RelationShipService
}

func (r *relationshipController) Initialize() {
}

func (r *relationshipController) Name() string {
	return "relationship"
}

func (r *relationshipController) SetupRoutes() {
	mw := JWT.Middleware()
	r.createLink = route.NewRouter("/api/relation", "POST")
	{
		r.createLink.SetMiddleWare(mw)
		r.createLink.SetHandler(r.createBindLink)
	}
	r.activeLink = route.NewRouter("/api/relation/:link", "GET")
	{
		r.activeLink.SetMiddleWare(mw)
		r.activeLink.SetHandler(r.linkUser)
	}
	r.deleteLink = route.NewRouter("/api/relation/", "DELETE")
	{
		r.deleteLink.SetMiddleWare(mw)
		r.deleteLink.SetHandler(r.unbindUser)
	}
	r.getRelation = route.NewRouter("/api/relation/list/:user", "GET")
	{
		r.getRelation.SetMiddleWare(mw)
		r.getRelation.SetHandler(r.getFriendList)
	}
	r.associateHistory = route.NewRouter("/api/relation/history/:user", "GET")
	{
		r.associateHistory.SetMiddleWare(mw)
		r.associateHistory.SetHandler(r.getHistory)
	}
}

func (r *relationshipController) GetRoutes() []Interface2.IRoute {
	return []Interface2.IRoute{r.createLink, r.activeLink, r.deleteLink, r.getRelation, r.associateHistory}
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
		UserID       int64 `json:"user_id,omitempty" form:"user_id"`
		RelationType int   `json:"relation_type,omitempty" form:"relation_type"`
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

	type RelationShipResponse struct {
		URL          string `json:"url"`
		RelationType int    `json:"relation_type"` // identify the relation type
	}

	var urlResp RelationShipResponse
	urlResp.URL = "/api/relation/" + link
	urlResp.RelationType = req.RelationType
	resp.Success(urlResp)
}

// @Title linkUser
// @description start to process the user's associate link jobs
func (r *relationshipController) linkUser(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	// get the user's associate link
	link := ctx.Param("link")
	if link == "" {
		resp.Error("invalid link")
		return
	}

	//todo: process the link

	type request struct {
		UserID int64 `json:"user_id,omitempty" form:"user_id"` // this id is the receiver's id
	}
	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error(err)
		resp.Error("invalid request")
		return
	}

	err := r.service.BindingTwoUser(link, req.UserID)
	if err != nil {
		resp.Error("failed to bind the user")
		return
	}

	resp.Success("success")

}

func (r *relationshipController) unbindUser(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type request struct {
		UserID   int64 `json:"user_id,omitempty" form:"user_id"`
		TargetID int64 `json:"target_id,omitempty" form:"target_id"`
	}

	var req request
	if err := ctx.ShouldBind(&req); err != nil {
		log.Error(err)
		resp.Error("invalid request")
		return
	}

	if !r.service.DisassociateUser(req.UserID, req.TargetID) {
		resp.Error("failed to disassociate the user")
		return
	}

	resp.Success("success")

}

func (r *relationshipController) getFriendList(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	user := ctx.Param("user")
	log.Info("user: ", user)
	if user == "" {
		resp.Error("invalid user")
		return
	}

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	lists := r.service.GetFriendList(claim.Id)

	resp.Success(lists)
}

func (r *relationshipController) getHistory(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	user := ctx.Param("user")
	if user == "" {
		resp.Error("invalid user")
		return
	}

	claim := getAuthObject(ctx)
	log.Debugf("uid: %d", claim.Id)
	if claim == nil {
		resp.Unauthorized("please login first")
		return
	}

	if claim.UserName != user {
		resp.Error("invalid request")
		return
	}

	lists := r.service.GetHistoryList(claim.Id)

	resp.Success(lists)
}

func NewRelationshipController(service services.RelationShipService) Interface2.IController {
	controller := &relationshipController{
		service: service,
	}
	return controller
}
