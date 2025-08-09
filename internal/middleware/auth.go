package middleware

import (
	"go-hospital-api/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware validates JWT token and injects staffID into Gin context
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		staffID, err := utils.VerifyTokenFromRequest(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "unauthorized: " + err.Error()})
			c.Abort()
			return
		}

		c.Set("staffID", staffID)
		c.Next()
	}
}
