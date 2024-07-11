package controller

import (
	"github.com/Thenecromance/OurStories/Interface"
	"github.com/Thenecromance/OurStories/application/services"
	"github.com/Thenecromance/OurStories/response"
	"github.com/Thenecromance/OurStories/route"
	"github.com/Thenecromance/OurStories/utility/log"
	"github.com/gin-gonic/gin"
)

type shopRoutes struct {
	shop Interface.IRoute
	cart Interface.IRoute
}

type shopControllerImpl struct {
	service services.ShopService

	//cart shop.ShopService
	shop services.ShopService

	routes shopRoutes
}

func (s *shopControllerImpl) Initialize() {
	//TODO implement me
	panic("implement me")
}

func (s *shopControllerImpl) Name() string {
	return "shopController"
}

func (s *shopControllerImpl) SetupRoutes() {
	//get shop items
	s.routes.shop = route.NewREST("/api/shop/")
	{
		s.routes.shop.SetHandler(s.getShopItems, s.addItemToShop, s.updateShopItem, s.deleteShopItem)
	}
	s.routes.cart = route.NewREST("/api/cart/:id") // this path
	{
		s.routes.cart.SetHandler(s.userGetCartData, s.userAddItemToCart, s.userRemoveItemFromCart, s.userCleanCart)
	}

}

func (s *shopControllerImpl) GetRoutes() []Interface.IRoute {
	//TODO implement me
	panic("implement me")
}

// getShopItems can publish all available items in the shop
func (s *shopControllerImpl) getShopItems(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// addItemToShop for adding new item to the shop when a user want to sell something
func (s *shopControllerImpl) addItemToShop(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// updateShopItem for updating the item in the shop (only the owner of the item can update the item)
func (s *shopControllerImpl) updateShopItem(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// deleteShopItem for deleting the item from the shop (only the owner of the item can delete the item)
func (s *shopControllerImpl) deleteShopItem(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

}

// userGetCartData for getting the cart data of the user
func (s *shopControllerImpl) userGetCartData(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type Param struct {
		UserId int64
		CartId int64
	}
}

func (s *shopControllerImpl) userAddItemToCart(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type Param struct {
		UserId int64 `json:"user_id,omitempty" form:"user_id"`
		//CartId int64 `json:"cart_id,omitempty" form:"cart_id"` // could be 0 or empty
		ItemId int64 `json:"item_id,omitempty" form:"item_id"` // item not null
		Count  int   `json:"count,omitempty" form:"count"`     // count of the item
	}
	var param Param
	if err := ctx.ShouldBind(&param); err != nil {
		resp.Error("invalid request")
		return
	}

	if param.ItemId == 0 || param.Count == 0 {
		resp.Error("invalid item id or count")
		return
	}

	//if param.CartId == 0 {
	//	// if cart id is not provided, a new cart will be created
	//	cartId, err := s.shop.CreateCart(param.UserId)
	//	if err != nil {
	//		resp.Error("failed to create cart")
	//		return
	//	}
	//	param.CartId = cartId
	//}

	cartId, err := s.shop.GetCart(param.UserId)
	if err != nil {
		log.Error(err)
		resp.Error("failed to get cart")
		return
	}

	err = s.shop.AddItemIntoCart(cartId, param.ItemId, param.Count)
	if err != nil {
		resp.Error("failed to add item into cart")
		return
	}
	resp.Success("done")
}

func (s *shopControllerImpl) userRemoveItemFromCart(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type Param struct {
		UserId int64
		//CartId int64
		ItemId int64
		Count  int
	}
	var param Param
	if err := ctx.ShouldBind(&param); err != nil {
		resp.Error("invalid request")
		return
	}

	//todo : check user id is same as it's jwt token

	cartId, err := s.shop.GetCart(param.UserId)

	// precheck the user's operation is valid or not
	if err != nil {
		log.Warn(err)
		resp.Error("failed to get cart")
		return
	}

	if param.ItemId == 0 || param.Count == 0 {
		log.Warn("received cart id is not same as the user's cart id")
		resp.Error("invalid cart id")
		return
	}

	// prepare to remove the items from the cart
	err = s.shop.RemoveItemFromCart(cartId, param.ItemId, param.Count)

}

func (s *shopControllerImpl) userCleanCart(ctx *gin.Context) {
	resp := response.New()
	defer resp.Send(ctx)

	type Param struct {
		UserId int64
	}

	var param Param
	if err := ctx.ShouldBind(&param); err != nil {
		resp.Error("invalid request")
		return
	}

	cartId, err := s.shop.GetCart(param.UserId)
	if err != nil {
		log.Error(err)
		resp.Error("failed to get cart")
		return
	}
	if err := s.shop.CleanCart(cartId); err != nil {
		log.Error(err)
		resp.Error("failed to clean cart")
		return
	}

	resp.Success("done")
}

func NewShopController() Interface.IController {
	return &shopControllerImpl{}
}
