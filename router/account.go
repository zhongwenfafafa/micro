package router

import (
	"github.com/gin-gonic/gin"
	"micro/middleware"

	"micro/service/apigw/account"
)

func Router()*gin.Engine {
	router := gin.Default()
	// 注册中间件
	router.Use(middleware.TraceIdGenerate())

	// 获取控制器实例
	accountController := account.Account{}
	//不需要验证就能访问的接口
	router.POST("/signup", accountController.SinUpHandler)

	return router
}