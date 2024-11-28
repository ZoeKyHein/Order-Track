package controllers

import (
	"github.com/gin-gonic/gin"
	"order-track/dto"
	"order-track/services"
	"order-track/utils"
)

// AddCustomer 新增客户
func AddCustomer(c *gin.Context) {
	req := dto.AddCustomerRequest{}

	// 绑定参数
	if err := c.ShouldBind(&req); err != nil {
		utils.Fail(c, nil, 1001, err.Error())
		return
	}

	// 调用服务
	customer, err := services.AddCustomer(req)
	if err != nil {
		utils.Fail(c, nil, 1002, err.Error())
	}

	// 返回数据
	utils.Success(c, customer, "新增客户成功")
}
