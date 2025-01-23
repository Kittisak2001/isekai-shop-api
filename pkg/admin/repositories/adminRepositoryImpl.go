package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
)

type (
	adminRepositoryImpl struct {
		db     databases.Database
	}
)

func NewAdminRepositoryImpl(db databases.Database) AdminRepository {
	return &adminRepositoryImpl{db}
}

func (r *adminRepositoryImpl) Creating(adminEntity *entities.Admin) (*entities.Admin, error) {
	if err := r.db.Connect().Create(adminEntity).Error; err != nil{
		return nil, err
	}
	return adminEntity, nil
}

func (r *adminRepositoryImpl) FindByID(adminID string) (*entities.Admin, error) {
	adminEntity := new(entities.Admin)
	if err := r.db.Connect().Where("id = ?", adminID).Take(adminEntity).Error; err != nil{
		return nil, err
	}
	return adminEntity, nil
}