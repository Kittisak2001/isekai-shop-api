package repositories

import "github.com/Kittisak2001/isekai-shop-api/entities"

type PlayerCoinRepository interface {
	CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string)(*entities.PlayerCoin, error)
}