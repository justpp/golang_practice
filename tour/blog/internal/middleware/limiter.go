package middleware

import (
	"giao/tour/blog/pkg/app"
	"giao/tour/blog/pkg/errcode"
	"giao/tour/blog/pkg/limiter"
	"github.com/gin-gonic/gin"
)

func RateLimiter(l limiter.Iface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				app.NewResponse(c).ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
