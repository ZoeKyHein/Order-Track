package services

import (
	"fmt"
	"order-track/dto"
	"order-track/models"
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
		return dto.CustomerDTO{}, fmt.Errorf("新增客户失败: %v", err)
	}

	return customer.TransferToDTO(), nil

}
