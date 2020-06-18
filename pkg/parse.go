package pkg

import (
	"github.com/gin-gonic/gin"
	"micro/defined"
	"micro/response"
	"net/http"
)

// 解析请求参数
func ParseRequest(c *gin.Context, request interface{}) error {
	err := c.ShouldBind(request)
	if err != nil {
		c.JSON(http.StatusOK,
			response.ErrorResponse{
				Code:    defined.VALIDATE_ERROR_CODE,
				Message: err.Error(),
			})
	}

	return err
}