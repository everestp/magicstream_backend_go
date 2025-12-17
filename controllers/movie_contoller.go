package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/everestp/magicstream_backend_go/database"
	"github.com/everestp/magicstream_backend_go/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

  var movieCollection *mongo.Collection =  database.OpenCollection("movies")
func GetMovies()  gin.HandlerFunc {
	return  func(c *gin.Context) {
	    ctx , cancel := context.WithTimeout(context.Background(), 100 * time.Second)
		defer  cancel()
		var movies []models.Movie
		cursor , err :=  movieCollection.Find(ctx, bson.M{})
		if  err != nil {
			c.JSON( http.StatusInternalServerError, gin.H{"error":"Failed to  fetch movie"})
		}
		//  for memory management close the connection
		defer cursor.Close(ctx)

		if err = cursor.All(ctx, &movies); err != nil{
			c.JSON( http.StatusInternalServerError, gin.H{"error":"Failed to  decode movie"})
		}
		c.JSON(http.StatusOK , movies)

	
	}
}


// contorller return  one movie
  func GetMovie()   gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx ,cancel :=context.WithTimeout(context.Background() , 100*time.Second)
		defer cancel()
		moviesID := c.Param("imdb_id")
		 if moviesID == ""{
			c.JSON( http.StatusBadGateway, gin.H{"error":"Movies ID Required"})
			return 
		 }
		 var movie models.Movie
		 err := movieCollection.FindOne(ctx , bson.M{"imdb_id":moviesID}).Decode(&movie)
		 if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error":"Movie not found"})
			return 
		 }
		 c.JSON(http.StatusOK, movie)
	}
  }
