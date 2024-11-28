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

// FetchAllCustomers 获取客户列表
func FetchAllCustomers(req dto.FetchCustomerRequest) ([]dto.CustomerDTOWithStats, int64, error) {

	// 设置分页默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}

	// 查询客户
	customers, total, err := models.FetchAllCustomers(req)
	if err != nil {
		return nil, 0, fmt.Errorf("获取客户列表失败: %v", err)
	}

	// 转换为DTO
	var customerDTOs []dto.CustomerDTOWithStats
	for _, customer := range customers {
		customerDTOs = append(customerDTOs, customer.TransferToDTO())
	}

	return customerDTOs, total, nil
}
