package controllers

import (
	"net/http"
	_custom "github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	_itemShopModel "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/model"
	_itemShopService "github.com/Kittisak2001/isekai-shop-api/pkg/itemShop/services"
	"github.com/labstack/echo/v4"
)

type itemShopControllerImpl struct {
	itemShopService _itemShopService.ItemShopService
}

func NewItemShopControllerImpl(itemShopService _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopService}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemFilter := new(_itemShopModel.ItemFilter)
	if err := _custom.NewEchoRequest(pctx).Bind(itemFilter); err != nil {
		return _custom.Error(pctx, http.StatusBadRequest, err.Error())
	}
	itemModelList, err := c.itemShopService.Listing(itemFilter)
	if err != nil {
		return _custom.Error(pctx, http.StatusInternalServerError, err.Error())
	}
	return pctx.JSON(http.StatusOK, itemModelList)
}