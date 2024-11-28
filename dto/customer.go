package dto

// AddCustomerRequest 新增客户请求
type AddCustomerRequest struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Remark  string `json:"remark"`
}

// CustomerDTO 客户数据传输对象
type CustomerDTO struct {
	ID        uint   `json:"id"`        // 客户ID
	Name      string `json:"name"`      // 姓名
	Phone     string `json:"phone"`     // 手机号
	Email     string `json:"email"`     // 邮箱
	Address   string `json:"address"`   // 地址
	Remark    string `json:"remark"`    // 备注
	CreatedAt string `json:"createdAt"` // 创建时间
	UpdatedAt string `json:"updatedAt"` // 更新时间
}
