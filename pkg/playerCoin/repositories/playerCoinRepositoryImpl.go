package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

type playerCoinRepositoryImpl struct {
	db databases.Database
}

func NewPlayerCoinRepositoryImpl(db databases.Database) PlayerCoinRepository {
	return &playerCoinRepositoryImpl{db}
}

func (r *playerCoinRepositoryImpl) CoinAdding(tx *gorm.DB,playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	conn := r.db.Connect()
	if tx != nil{
		conn = tx
	}
	if err := conn.Model(&entities.PlayerCoin{}).Create(playerCoinEntity).Error; err != nil {
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