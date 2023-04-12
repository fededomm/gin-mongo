// go:build (pippo)
// +build pippo

package routes

import (
	"gin-mongo/configuration"
	"gin-mongo/docs"
	"gin-mongo/middleware"
	"gin-mongo/assets"
	"html/template"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)
func Init(host *configuration.RouterConf,
	db *mongo.Client,
	servicename string,
	) {

	docs.SwaggerInfo.BasePath = "/api/v1/"
	
	docs.SwaggerInfo.Host = "192.168.3.109:8085"
	docs.SwaggerInfo.Description = "Test API for Ordini"
	docs.SwaggerInfo.Title = "Ordini API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	var routes Routes
	routes.DB = db
	router := gin.Default()
	
	t, err := assets.LoadTemplate()
	if err != nil {
		panic(err.Error())
	}

	router.SetHTMLTemplate(t)
	router.Use(middleware.Middleware())
	router.Use(otelgin.Middleware("gin-mongo-middle"))

	v1 := router.Group("/api/v1")
	{
		Gestionale := v1.Group("/gest")
		{
			Gestionale.GET("", routes.GetOrdini)
			Gestionale.POST("", routes.PostOrdini)
			Gestionale.PUT(":numeroOrdine", routes.UpdateOrdine)
			Gestionale.DELETE(":numeroOrdine", routes.DeleteOrdine)
		}
	}
	router.GET("/health", func(c *gin.Context) {
		c.HTML(200, "healthcheck.tmpl", gin.H{
			"title": "SERVER UP"})
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(host.Router) // NON DIMENTICARSI IL SERVE!!!!!
}
