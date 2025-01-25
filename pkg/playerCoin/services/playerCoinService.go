package services

import "github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"

type PlayerCoinService interface {
	CoinAdding(coinAddingReq *model.CoinAddingReq) (*model.PlayerCoin, error)
	Showing(playerID string) (*model.PlayerCoinShowing)
}