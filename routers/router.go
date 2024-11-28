package routes

import (
	"github.com/gin-gonic/gin"
	"order-track/controllers"
	"order-track/middleware"
	"order-track/utils"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.GinZapLogger(utils.Log))
	customer := router.Group("/customer")
	{
		customer.POST("", controllers.AddCustomer) // 新增用户
	}

	return router
}
