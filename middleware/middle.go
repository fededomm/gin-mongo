package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		log.Print("Response Headers: ")
		log.Println(c.Request.Header)
		
		c.Next()
	}
}
