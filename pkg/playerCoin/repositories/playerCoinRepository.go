package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type PlayerCoinRepository interface {
	CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error)
	Showing(playerID string) (*entities.PlayerCoin, error)
}