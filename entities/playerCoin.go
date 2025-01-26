package entities

import (
	"time"
	"github.com/Kittisak2001/isekai-shop-api/pkg/playerCoin/model"
)

type (
	PlayerCoin struct {
		ID        uint64    `gorm:"primaryKey;autoIncrement;"`
		PlayerID  string    `gorm:"type:varchar(64);not null;"`
		Amount    int64     `gorm:"not null;"`
		CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
		Coin      int64     `gorm:"->" json:"coin"`
	}
)

func (e *PlayerCoin) ToPlayerCoinModel() *model.PlayerCoin {
	return &model.PlayerCoin{
		ID:        e.ID,
		PlayerID:  e.PlayerID,
		Amount:    e.Amount,
		CreatedAt: e.CreatedAt,
	}
}

func (e *PlayerCoin) ToPlayerCoinShowingModel() *model.PlayerCoinShowing {
	return &model.PlayerCoinShowing{
		PlayerID: e.PlayerID,
		Coin: e.Coin,
	}
}