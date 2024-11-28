package services

import (
	"fmt"
	"go.uber.org/zap"
	"order-track/dto"
	"order-track/models"
	"order-track/utils"
)

// AddCustomer 新增客户
func AddCustomer(req dto.AddCustomerRequest) (dto.CustomerDTO, error) {

	customer := models.Customer{
		Name:    req.Name,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		Remark:  req.Remark,
	}

	if err := customer.AddCustomer(); err != nil {
		utils.Log.Error("新增客户失败", zap.String("name", req.Name), zap.String("phone", req.Phone), zap.Error(err))
		return dto.CustomerDTO{}, fmt.Errorf("新增客户失败: %v", err)
	}

	utils.Log.Info("新增客户成功",
		zap.String("name", req.Name),
		zap.String("phone", req.Phone),
	)
	return customer.TransferToDTO(), nil

}
