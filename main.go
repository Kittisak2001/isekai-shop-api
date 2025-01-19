package main

import (
	"github.com/Kittisak2001/isekai-shop-api/config"
	"github.com/Kittisak2001/isekai-shop-api/databases"
	"github.com/Kittisak2001/isekai-shop-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db.ConnectionGetting())

	server.Start()
}