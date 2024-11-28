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
		utils.Fail(c, nil, utils.ErrCodeBindError, err.Error())
		return
	}

	// 调用服务
	customer, err := services.AddCustomer(req)
	if err != nil {
		utils.Fail(c, nil, utils.ErrCodeServiceError, err.Error())
		return
	}

	// 返回数据
	utils.Success(c, customer, "新增客户成功")
}

// FetchAllCustomers 获取客户列表
func FetchAllCustomers(c *gin.Context) {
	req := dto.FetchCustomerRequest{}

	// 绑定参数
	if err := c.ShouldBind(&req); err != nil {
		utils.Fail(c, nil, utils.ErrCodeBindError, err.Error())
		return
	}

	// 调用服务
	customers, total, err := services.FetchAllCustomers(req)
	if err != nil {
		utils.Fail(c, nil, utils.ErrCodeServiceError, err.Error())
		return
	}

	// 返回数据
	utils.Success(c, gin.H{
		"list":  customers,
		"total": total,
	}, "获取客户列表成功")
}
