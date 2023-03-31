package routes

import (
	"gin-mongo/configuration"
	"gin-mongo/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(host *configuration.RouterConf,
	db *mongo.Client,
	servicename string) {

	docs.SwaggerInfo.BasePath = "/api/v1/"
	docs.SwaggerInfo.Host = "192.168.3.109:8085"
	docs.SwaggerInfo.Description = "Test API for Ordini"
	docs.SwaggerInfo.Title = "Ordini API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	var routes Routes
	routes.DB = db
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		Gestionale := v1.Group("/gest")
		{
			Gestionale.GET("", routes.GetOrdini)
			Gestionale.POST("", routes.PostOrdini)
			Gestionale.PUT(":numeroOrdine", routes.UpdateOrdine)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(host.Router) // NON DIMENTICARSI IL SERVE!!!!!
}
