package main

import (
	"fmt"

	routes "github.com/everestp/magicstream_backend_go/roues"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello magic stream")

	})
    routes.SetupProtectedRoutes(router, client *mongo.Client)
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Failed to start server", err)
	}
	fmt.Println("This is the first  go  code")
}
