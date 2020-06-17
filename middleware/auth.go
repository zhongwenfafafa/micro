package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if len(token) == 0 {
			c.Abort()
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		c.Next()
	}
}