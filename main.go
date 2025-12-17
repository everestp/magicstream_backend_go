package main

import (
	"fmt"

	controller "github.com/everestp/magicstream_backend_go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello magic stream")

	})
	router.GET("/movies", controller.GetMovies())
	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
	fmt.Println("This is the first  go  code")
}
