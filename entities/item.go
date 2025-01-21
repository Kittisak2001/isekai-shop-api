package entities

import "time"

type (
	Item struct {
		ID          uint64    `gorm:"primaryKey;autoIncrement;"`
		AdminID     *string   `gorm:"type:varchar(64);"`
		Name        string    `gorm:"type:varchar(64);unique;not null;"`
		Description string    `gorm:"type:varchar(128);not null;"`
		Price       int       `gorm:"not null;default:0;"`
		Picture     string    `gorm:"not null;"`
		IsArchive   bool      `gorm:"not null;default:false;"`
		CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
		UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
	}
)