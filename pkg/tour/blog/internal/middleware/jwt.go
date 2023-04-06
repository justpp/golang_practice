package middleware

import (
	"giao/pkg/tour/blog/pkg/app"
	"giao/pkg/tour/blog/pkg/errcode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() func(c *gin.Context) {
	return func(c *gin.Context) {
		var (
			token string
			code  = errcode.Success
		)
		tokenName := "token"
		query, b := c.GetQuery(tokenName)
		if b {
			token = query
		} else {
			token = c.GetHeader(tokenName)
		}
		if token == "" {
			code = errcode.InvalidParams
		} else {
			_, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = errcode.UnauthorizedTokenTimeout
				default:
					code = errcode.UnauthorizedTokenError
				}
			}
		}

		if code != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(code)
			c.Abort()
			return
		}

		c.Next()
	}
}
