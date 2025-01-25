package services

import (
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/exception"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/repositories"
	"github.com/labstack/echo/v4"
)

type playerCoinServiceImpl struct {
	playerCoinRepository repositories.PlayerCoinRepository
	logger               echo.Logger
}

func NewPlayerCoinServiceImpl(playerCoinRepository repositories.PlayerCoinRepository, logger echo.Logger) PlayerCoinService {
	return &playerCoinServiceImpl{playerCoinRepository, logger}
}

func (s *playerCoinServiceImpl) CoinAdding(coinAddingReq *model.CoinAddingReq) (*model.PlayerCoin, error) {
	playerCoinEntity := &entities.PlayerCoin{
		PlayerID: coinAddingReq.PlayerID,
		Amount:   coinAddingReq.Amount,
	}
	playerCoinEntity, err := s.playerCoinRepository.CoinAdding(playerCoinEntity)
	if err != nil {
		s.logger.Errorf("Error to coin adding: %s", err.Error())
		return nil, &exception.CoinAdding{}
	}
	return playerCoinEntity.ToPlayerCoinModel(), nil
}

func (s *playerCoinServiceImpl) Showing(playerID string) *model.PlayerCoinShowing {
	playerCoinEntity, err := s.playerCoinRepository.Showing(playerID)
	if err != nil {
		return &model.PlayerCoinShowing{
			PlayerID: playerID,
			Coin:     0,
		}
	}
	return playerCoinEntity.ToPlayerCoinShowingModel()
}