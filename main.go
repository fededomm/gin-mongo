package main

import (
	_ "embed"
	"fmt"
	db "gin-mongo/database"
	"gin-mongo/routes"
	"log"
)

//go:embed  banner.txt
var banner []byte

func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := *db.ConnectDB(&cfg.App.Database)

	fmt.Println(string(banner))
	log.Println("Connection String: " + cfg.App.Database.ConnectionString)
	routes.Init(&cfg.App.Router, &mongoClient, cfg.App.ServiceName)
}
