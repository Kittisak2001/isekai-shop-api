package repository

import "github.com/Kittisak2001/isekai-shop-api/entities"

type (
	PlayerRepository interface {
		Creating(playerEntity *entities.Player) (*entities.Player, error)
		FindByID(playerID string) (*entities.Player, error)
	}
)