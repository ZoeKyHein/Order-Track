package models

import (
	"gorm.io/gorm"
	"order-track/config"
	"order-track/dto"
	"order-track/utils"
	"time"
)

// Customer 客户模型
type Customer struct {
	ID        uint           `gorm:"primaryKey"`              // 客户ID
	Name      string         `gorm:"size:255;not null"`       // 姓名
	Phone     string         `gorm:"size:20;unique;not null"` // 手机号
	Email     string         `gorm:"size:255"`                // 邮箱
	Address   string         `gorm:"type:text"`               // 地址
	Remark    string         `gorm:"type:text"`               // 备注
	CreatedAt time.Time      `gorm:"autoCreateTime"`          // 创建时间
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`          // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"`                   // 删除时间
}

// AddCustomer 新增客户
func (c *Customer) AddCustomer() error {
	return config.DB.Create(&c).Error
}

func (c *Customer) TransferToDTO() dto.CustomerDTO {
	return dto.CustomerDTO{
		ID:        c.ID,
		Name:      c.Name,
		Phone:     c.Phone,
		Email:     c.Email,
		Address:   c.Address,
		Remark:    c.Remark,
		CreatedAt: c.CreatedAt.Format(utils.TimeFormat),
		UpdatedAt: c.UpdatedAt.Format(utils.TimeFormat),
	}
}

// CustomerWithStats 客户及统计信息模型
type CustomerWithStats struct {
	Customer
	OrderCount int     `gorm:"order_count"` // 客户订单数
	TotalSales float64 `gorm:"total_sales"` // 客户累计销售额
}

func (cws *CustomerWithStats) TransferToDTO() dto.CustomerDTOWithStats {
	return dto.CustomerDTOWithStats{
		CustomerDTO: dto.CustomerDTO{
			ID:        cws.ID,
			Name:      cws.Name,
			Phone:     cws.Phone,
			Email:     cws.Email,
			Address:   cws.Address,
			Remark:    cws.Remark,
			CreatedAt: cws.CreatedAt.Format(utils.TimeFormat),
			UpdatedAt: cws.UpdatedAt.Format(utils.TimeFormat),
		},
		OrderCount:  cws.OrderCount,
		AmountCount: cws.TotalSales,
	}
}

// FetchAllCustomers 获取客户列表
func FetchAllCustomers(req dto.FetchCustomerRequest) ([]CustomerWithStats, int64, error) {
	var customers []CustomerWithStats
	var total int64

	db := config.DB.Model(&Customer{}).Select(`customers.*,
        COUNT(orders.id) AS order_count,
        IFNULL(SUM(orders.amount), 0) AS total_sales`).
		Joins("LEFT JOIN orders ON customers.id = orders.customer_id").
		Group("customers.id")

	if req.KeyWords != "" {
		db = db.Where(`customers.name LIKE ?`, "%"+req.KeyWords+"%")
	}

	if req.Phone != "" {
		db = db.Where(`customers.phone = ?`, req.Phone)
	}

	limit := req.PageSize
	offset := (req.Page - 1) * req.PageSize

	err := db.Count(&total).Limit(limit).Offset(offset).Find(&customers).Error
	if err != nil {
		return nil, 0, err
	}
	return customers, total, nil
}
