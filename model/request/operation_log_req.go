package request

// OperationLogListReq 操作日志请求结构体
type OperationLogListReq struct {
	Username string `json:"username" form:"username"`
	Ip       string `json:"ip" form:"ip"`
	Path     string `json:"path" form:"path"`
	Status   int    `json:"status" form:"status"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// OperationLogDeleteReq 批量删除操作日志结构体
type OperationLogDeleteReq struct {
	OperationLogIds []uint `json:"operationLogIds" validate:"required"`
}
