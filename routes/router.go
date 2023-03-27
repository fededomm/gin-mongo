package routes

import (
	"gin-mongo/configuration"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(host *configuration.RouterConf,
	db *mongo.Client,
	servicename string) {

	var routes Routes
	routes.DB = db
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		Gestionale := v1.Group("/gest")
		{
			Gestionale.GET("", routes.GetOrdini)
			Gestionale.POST("", routes.PostOrdini)
		}
	}
	router.Run(host.Router) // NON DIMENTICARSI IL SERVE!!!!!
}
