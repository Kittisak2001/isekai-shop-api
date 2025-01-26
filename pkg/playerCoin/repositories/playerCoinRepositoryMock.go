package repositories

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type playerCoinRepositoryMock struct {
	mock.Mock
}

func NewPlayerCoinRepositoryMock() *playerCoinRepositoryMock {
	return &playerCoinRepositoryMock{}
}

func (m *playerCoinRepositoryMock) CoinAdding(tx *gorm.DB, playerCoinEntity *entities.PlayerCoin) (*entities.PlayerCoin, error) {
	args := m.Called(tx, playerCoinEntity)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}
func (m *playerCoinRepositoryMock) Showing(playerID string) (*entities.PlayerCoin, error) {
	args := m.Called(playerID)
	return args.Get(0).(*entities.PlayerCoin), args.Error(1)
}