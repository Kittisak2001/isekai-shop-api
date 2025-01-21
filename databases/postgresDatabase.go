package databases

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"github.com/Kittisak2001/isekai-shop-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), 
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		)

		dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
			conf.Host, conf.Port, conf.User, conf.Password, conf.DBName, conf.SSLMode, conf.Schema,
		)
		conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: newLogger,
		})
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