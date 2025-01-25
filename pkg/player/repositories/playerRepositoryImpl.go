package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
)

type (
	playerRepositoryImpl struct {
		db databases.Database
	}
)

func NewPlayerRepositoryImpl(db databases.Database) PlayerRepository {
	return &playerRepositoryImpl{db}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	if err := r.db.Connect().Create(playerEntity).Error; err != nil {
		return nil, err
	}
	return playerEntity, nil
}
func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	playerEntity := new(entities.Player)
	if err := r.db.Connect().Where("id = ?", playerID).Take(playerEntity).Error; err != nil {
		return nil, err
	}
	return playerEntity, nil
}