package middleware

import "github.com/gin-gonic/gin"

func ApiInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("api_name", "blob_service")
		c.Set("api_version", "1.0.0")
		c.Next()
	}
}
