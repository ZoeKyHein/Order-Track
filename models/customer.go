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
