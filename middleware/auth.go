package middleware

import (
	"fmt"
	"pplx2api/config"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		fmt.Println("[DEBUG] Authorization Header:", authHeader)

		if authHeader != "" {
			key := strings.TrimPrefix(authHeader, "Bearer ")
			fmt.Println("[DEBUG] Parsed API Key from header:", key)
			fmt.Println("[DEBUG] Expected API Key (from env/config):", config.ConfigInstance.APIKey)

			if key != config.ConfigInstance.APIKey {
				c.JSON(401, gin.H{
					"error": "Invalid API key",
				})
				c.Abort()
				return
			}

			c.Next()
			return
		}

		// No Authorization header at all
		c.JSON(401, gin.H{
			"error": "Missing or invalid Authorization header",
		})
		c.Abort()
	}
}
