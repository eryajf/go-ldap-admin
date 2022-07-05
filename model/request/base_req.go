package request

// BaseChangePwdReq 修改密码结构体
type BaseChangePwdReq struct {
	Mail string `json:"mail" validate:"required,min=0,max=100"`
}

// BaseDashboardReq  系统首页展示数据结构体
type BaseDashboardReq struct {
}

// GetPasswdReq
type GetPasswdReq struct {
	Passwd string `json:"passwd" form:"passwd" validate:"required"`
}
