package routes

import (
	"github.com/gin-gonic/gin"
	"order-track/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	customer := router.Group("/customer")
	{
		customer.POST("", controllers.AddCustomer) // 新增用户
	}

	return router
}
