package middleware

import (
	"github.com/china-golang/go-apiservice/handler"
	"github.com/china-golang/go-apiservice/pkg/errno"
	"github.com/china-golang/go-apiservice/pkg/token"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the json web token.
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
