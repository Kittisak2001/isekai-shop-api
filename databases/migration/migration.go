package main

import (
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	tx := db.ConnectionGetting().Begin()
	defer tx.Rollback()
	
	if err := tx.Debug().AutoMigrate(&entities.Item{}); err != nil{
		panic(err)
	}
	// playerCoinMigration(tx)
	// inventoryMigration(tx)
	// purchaseHistoryMigration(tx)
	// playerMigration(tx)
	// adminMigration(tx)
	// itemMigration(tx)
	if err :=	tx.Commit().Error; err != nil{
		panic(err)
	}
}

func playerMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Admin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Item{})
}

func playerCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PlayerCoin{})
}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Inventory{})
}

func purchaseHistoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.PurchaseHistory{})
}
