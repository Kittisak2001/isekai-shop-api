package controllers

import (
	"net/http"

	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/inventory/services"
	"github.com/Kittisak2001/isekai-shop-api/pkg/oAuth2/exception"
	"github.com/labstack/echo/v4"
)

type inventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(inventoryService services.InventoryService) *inventoryController {
	return &inventoryController{inventoryService: inventoryService}
}

func (c *inventoryController) Listing(pctx echo.Context) error {
	playerID, ok := pctx.Get("playerID").(string)
	if !ok || playerID == "" {
		return custom.Error(pctx, http.StatusBadRequest, &exception.PlayerNotFound{})
	}
	inventoryListing, err := c.inventoryService.Listing(playerID)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, inventoryListing)
}

func (c *inventoryController) Selling(pctx echo.Context) error {
	return nil
}