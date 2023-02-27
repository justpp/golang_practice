package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancelFunc := context.WithTimeout(c, t)
		defer cancelFunc()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
