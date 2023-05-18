package main

import (
	"context"
	_ "embed"
	"fmt"
	db "gin-mongo/database"
	routes "gin-mongo/routes"
	observability "gin-mongo/telemetry"
	"log"
)

//go:embed  banner.txt
var banner []byte

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host	127.0.0.1:8085
func main() {
	cfg, err := ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	mongoClient := *db.ConnectDB(&cfg.App.Database)
	if cfg.App.Observability.Enable {
		log.Println("***OPENTELEMETRY ON***")
		log.Println("Collector Endpoint: " + cfg.App.Observability.Endpoint)
		trace, err := observability.InitTracer(&cfg.App.Observability)
		if err != nil {
			log.Fatal("impossibile inizializzare il Tracer", err)
		}
		defer trace(context.Background())
	} else {
		log.Println("***OPENTELEMETRY OFF***")
	}
	fmt.Println(string(banner))
	log.Println("Connection String: " + cfg.App.Database.ConnectionString)
	routes.Init(&cfg.App.Router, &mongoClient, cfg.App.ServiceName)
}
