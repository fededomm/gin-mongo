package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		latency := time.Since(t)
		log.Println("inizio chiamata ")
		c.Next()
		log.Println("termino la chiamata in: " + latency.String())
	}
}
