package controllers

import (
	"net/http"
	"strconv"

	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/model"
	_itemManagingService "github.com/Kittisak2001/isekai-shop-api/pkg/itemManaging/services"
	"github.com/labstack/echo/v4"
)

type (
	itemManagingController struct {
		itemManagingService _itemManagingService.ItemManagingService
	}
)

func NewItemManagingController(itemManagingService _itemManagingService.ItemManagingService) *itemManagingController {
	return &itemManagingController{itemManagingService}
}

func (c *itemManagingController) Creating(pctx echo.Context) error {
	itemCreatingReq := new(model.ItemCreatingReq)
	if err := custom.NewEchoRequest(pctx).Bind(itemCreatingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	item, err := c.itemManagingService.Creating(itemCreatingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusCreated, item)
}

func (c *itemManagingController) Editing(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	itemEditingReq := new(model.ItemEditingReq)
	if err := custom.NewEchoRequest(pctx).Bind(itemEditingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	item, err := c.itemManagingService.Editing(itemID, itemEditingReq)
	if err != nil {
		exception := &exception.ItemNotfound{ItemID: *itemID}
		if err.Error() == exception.Error() {
			return custom.Error(pctx, http.StatusNotFound, err)
		}
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, item)
}

func (c *itemManagingController) Archiving(pctx echo.Context) error {
	itemID, err := c.getItemID(pctx)
	if err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	if err := c.itemManagingService.Archiving(itemID); err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.NoContent(http.StatusNoContent)
}

func (c *itemManagingController) getItemID(pctx echo.Context) (*uint64, error) {
	itemID := pctx.Param("itemID")
	itemIDUint64, err := strconv.ParseUint(itemID, 10, 64)
	if err != nil {
		return nil, err
	}
	return &itemIDUint64, nil
}
