package routes

import (
	 controller "github.com/everestp/magicstream_backend_go/controllers"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/v2/mongo"
)

func SetupUnProtectedRoutes(router *gin.Engine) {

	router.GET("/movies", controller.GetMovies())
	router.POST("/register", controller.RegisterUser())
	router.POST("/login", controller.LoginUser())
	// router.POST("/logout", controller.LogoutHandler(client))
	// router.GET("/genres", controller.GetGenres(client))
	// router.POST("/refresh", controller.RefreshTokenHandler(client))
}