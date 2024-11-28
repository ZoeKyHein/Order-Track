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

// FetchCustomerRequest 查询客户请求
type FetchCustomerRequest struct {
	Page     int    `json:"pageNum" form:"pageNum"`   // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页数量
	KeyWords string `json:"keyWords" form:"keyWords"` // 关键词
	Phone    string `json:"phone" form:"phone"`       // 手机号
}

// FetchCustomerResponse 查询客户响应
type FetchCustomerResponse struct {
	Total    int           `json:"total"`    // 总数
	Page     int           `json:"page"`     // 页码
	PageSize int           `json:"pageSize"` // 每页数量
	List     []CustomerDTO `json:"list"`     // 列表
}

type CustomerDTOWithStats struct {
	CustomerDTO
	OrderCount  int     `json:"orderCount"`  // 订单数量
	AmountCount float64 `json:"amountCount"` // 订单金额总量
}
