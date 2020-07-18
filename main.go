package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.Default()
	router.Use(TlsHandler())

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{"message":"pong"})
	})
	router.RunTLS(":35000", "./server.pem", "./server.key")
}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:35000",
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}
