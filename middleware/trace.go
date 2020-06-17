package middleware

import (
	"github.com/gin-gonic/gin"
	"micro/pkg"
)

// gin上下文携带全局唯一trace id
func TraceIdGenerate() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 上下文添加traceId
		pkg.GinContextTraceID(c)

		c.Next()
	}
}