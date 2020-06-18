package pkg

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"micro/defined"
)

func GinContextTraceID(c *gin.Context) {
	c.Set(defined.TRACE_KEY, GenerateUUID())
}

// 生成唯一uuid
func GenerateUUID() string {
	u, err :=  uuid.NewRandom()
	if err != nil {
		Logger.Warn("generate uuid error", zap.String("warn", err.Error()))
		u = uuid.New()
	}

	return u.String()
}