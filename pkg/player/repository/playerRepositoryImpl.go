package repository

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/labstack/echo/v4"
)

type (
	playerRepositoryImpl struct {
		db     databases.Database
		logger echo.Logger
	}
)

func NewPlayerRepositoryImpl(db databases.Database, logger echo.Logger) PlayerRepository {
	return &playerRepositoryImpl{db, logger}
}

func (r *playerRepositoryImpl) Creating(playerEntity *entities.Player) (*entities.Player, error) {
	if err := r.db.Connect().Create(playerEntity).Error; err != nil{
		return nil, err
	}
	return playerEntity, nil
}
func (r *playerRepositoryImpl) FindByID(playerID string) (*entities.Player, error) {
	playerEntity := new(entities.Player)
	if err := r.db.Connect().Where("id = ?", playerID).Take(playerEntity).Error; err != nil{
		return nil, err
	}
	return nil, nil
}