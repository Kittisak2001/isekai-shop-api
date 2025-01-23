package controllers

import (
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/services"
)

type playerCoinController struct{
	playerCoinService services.PlayerCoinService
}

func NewPlayerCoinController(playerCoinService services.PlayerCoinService) *playerCoinController{
	return &playerCoinController{playerCoinService}
}