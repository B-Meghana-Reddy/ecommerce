package middleware
// This file protects your APIs by checking whether the request has a valid login token before allowing access.
import (
	"net/http"

	token "github.com/B-Meghana-Reddy/ecommerce/tokens"

	"github.com/gin-gonic/gin"
)
// Returns a Gin middleware function that runs before protected routes.
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("uid", claims.Uid)
		c.Next()
	}
}