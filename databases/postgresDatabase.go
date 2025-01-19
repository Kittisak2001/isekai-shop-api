package databases

import (
	"fmt"
	"log"
	"sync"
	"github.com/Kittisak2001/isekai-shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	*gorm.DB
}

var (
	postgresDatabaseInstance *postgresDatabase
	once                     sync.Once
)

func NewPostgresDatabase(conf *config.DatabaseCfg) Database {
	once.Do(func() {
		dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode, conf.Schema,
		)
		conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		log.Printf("Connected to database %s", conf.DBName)
		postgresDatabaseInstance = &postgresDatabase{conn}
	})
	return postgresDatabaseInstance
}

func (db *postgresDatabase) ConnectionGetting() *gorm.DB {
	return db.DB
}