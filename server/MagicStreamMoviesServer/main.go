package main

import (
	"fmt"

	controller "github.com/Santhosh785/MagicStream/server/MagicStreamMoviesServer/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello, MagicStream")
	})
	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8081"); err != nil {
		fmt.Print("failed to start", err)
	}
}
