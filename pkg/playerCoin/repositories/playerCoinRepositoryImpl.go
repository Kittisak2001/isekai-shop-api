package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
)

type playerCoinRepositoryImpl struct {
	db databases.Database
}

func NewPlayerCoinRepositoryImpl(db databases.Database) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db}
}

func (r *playerCoinRepositoryImpl) CoinAdding(playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	if err := r.db.Connect().Create(playerCoinEntity).Error; err != nil {
		return nil, err
	}
	return playerCoinEntity, nil
}

func (r *playerCoinRepositoryImpl) Showing(playerID string) (*entities.PlayerCoin, error) {
	playerEntity := new(entities.PlayerCoin)
	if err := r.db.Connect().Select("player_id, sum(amount) as coin").Where("player_id = ?", playerID).Group("player_id").Take(&playerEntity).Error; err != nil {
		return nil, err
	}
	return playerEntity, nil
}