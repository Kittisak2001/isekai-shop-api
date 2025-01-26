package controllers

import (
	"net/http"
	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	_custom "github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) *itemShopControllerImpl {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)
	if err := _custom.NewEchoRequest(pctx).Bind(itemFilter); err != nil {
		return _custom.Error(pctx, http.StatusBadRequest, err)
	}
	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return _custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, itemModelList)
}

func (c *itemShopControllerImpl) Buying(pctx echo.Context) error {
	playerID, ok := pctx.Get("playerID").(string)
	if !ok || playerID == "" {
		return custom.Error(pctx, http.StatusBadRequest, &exception.PlayerNotFound{})
	}

	req := custom.NewEchoRequest(pctx)
	buyingReq := new(model.BuyingReq)
	if err := req.Bind(buyingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	buyingReq.PlayerID = playerID
	playerCoin, err := c.itemShopService.Buying(buyingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, playerCoin)
}

func (c *itemShopControllerImpl) Selling(pctx echo.Context) error {
	playerID, ok := pctx.Get("playerID").(string)
	if !ok || playerID == "" {
		return custom.Error(pctx, http.StatusBadRequest, &exception.PlayerNotFound{})
	}

	req := custom.NewEchoRequest(pctx)
	sellingReq := new(model.SellingReq)
	if err := req.Bind(sellingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	sellingReq.PlayerID = playerID
	playerCoin, err := c.itemShopService.Selling(sellingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, playerCoin)
}

