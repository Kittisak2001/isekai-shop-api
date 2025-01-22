package repository

import "github.com/Kittisak2001/isekai-shop-api/entities"

type (
	AdminRepository interface {
		Creating(adminEntity *entities.Admin) (*entities.Admin, error)
		FindByID(adminID string) (*entities.Admin, error)
	}
)