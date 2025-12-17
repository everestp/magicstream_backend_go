package routes

import (
	"github.com/everestp/magicstream_backend_go/middleware"
	 controller "github.com/everestp/magicstream_backend_go/controllers"
	"github.com/gin-gonic/gin"
	
)

func SetupProtectedRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleWare())

	router.GET("/movie/:imdb_id", controller.GetMovie())
	router.POST("/addmovie", controller.AddMovie())
	// router.GET("/recommendedmovies", controller.GetRecommendedMovies(client))
	// router.PATCH("/updatereview/:imdb_id", controller.AdminReviewUpdate(client))
}