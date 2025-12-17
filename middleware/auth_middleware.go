package middleware

import (
	"net/http"

	"github.com/everestp/magicstream_backend_go/utils"
	"github.com/gin-gonic/gin"
)





func AuthMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		token , err := utils.GetAccessToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort() 
		}
		if token == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error":"No Token Provided"} )
			c.Abort()
		}
		claims ,err := utils.ValidateRefreshToken(token)
		if err !=nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error":"No Token Provided"} )
			c.Abort()
			return 
		}
		c.Set("userID", claims.UserId)
		c.Set("role", claims.Role)
		c.Next()
	}
}