package controllers

import (
	"net/http"
	"github.com/Kittisak2001/isekai-shop-api/pkg/custom"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/services"
	"github.com/labstack/echo/v4"
)

type playerCoinController struct {
	playerCoinService services.PlayerCoinService
}

func NewPlayerCoinController(playerCoinService services.PlayerCoinService) *playerCoinController {
	return &playerCoinController{playerCoinService}
}

func (c *playerCoinController) CoinAdding(pctx echo.Context) error {
	playerID, ok := pctx.Get("playerID").(string)
	if !ok || playerID == "" {
		return custom.Error(pctx, http.StatusBadRequest, &exception.PlayerNotFound{})
	}
	coinAddingReq := new(model.CoinAddingReq)
	req := custom.NewEchoRequest(pctx)
	if err := req.Bind(coinAddingReq); err != nil {
		return custom.Error(pctx, http.StatusBadRequest, err)
	}
	coinAddingReq.PlayerID = playerID
	playerCoin, err := c.playerCoinService.CoinAdding(coinAddingReq)
	if err != nil {
		return custom.Error(pctx, http.StatusInternalServerError, err)
	}
	return pctx.JSON(http.StatusOK, playerCoin)
}

func (c *playerCoinController) Showing(pctx echo.Context) error {
	playerID, ok := pctx.Get("playerID").(string)
	if !ok || playerID == "" {
		return custom.Error(pctx, http.StatusBadRequest, &exception.PlayerNotFound{})
	}
	return pctx.JSON(http.StatusOK, c.playerCoinService.Showing(playerID))
}