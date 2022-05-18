package request

// RoleAddReq 添加资源结构体
type RoleAddReq struct {
	Name    string `json:"name" validate:"required,min=1,max=20"`
	Keyword string `json:"keyword" validate:"required,min=1,max=20"`
	Remark  string `json:"remark" validate:"min=0,max=100"`
	Status  uint   `json:"status" validate:"oneof=1 2"`
	Sort    uint   `json:"sort" validate:"gte=1,lte=999"`
}

// RoleListReq 列表结构体
type RoleListReq struct {
	Name     string `json:"name" form:"name"`
	Keyword  string `json:"keyword" form:"keyword"`
	Status   uint   `json:"status" form:"status"`
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// RoleUpdateReq 更新资源结构体
type RoleUpdateReq struct {
	ID      uint   `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required,min=1,max=20"`
	Keyword string `json:"keyword" validate:"required,min=1,max=20"`
	Remark  string `json:"remark" validate:"min=0,max=100"`
	Status  uint   `json:"status" validate:"oneof=1 2"`
	Sort    uint   `json:"sort" validate:"gte=1,lte=999"`
}

// RoleDeleteReq 删除资源结构体
type RoleDeleteReq struct {
	RoleIds []uint `json:"roleIds" validate:"required"`
}

// RoleGetTreeReq 获取资源树结构体
type RoleGetTreeReq struct {
}

// RoleGetMenuListReq 获取角色菜单列表结构体
type RoleGetMenuListReq struct {
	RoleID uint `json:"roleId" form:"roleId" validate:"required"`
}

// RoleGetApiListReq 获取角色接口列表结构体
type RoleGetApiListReq struct {
	RoleID uint `json:"roleId" form:"roleId" validate:"required"`
}

// RoleUpdateMenusReq 更新角色菜单结构体
type RoleUpdateMenusReq struct {
	RoleID  uint   `json:"roleId" validate:"required"`
	MenuIds []uint `json:"menuIds" validate:"required"`
}

// RoleUpdateApisReq 更新角色接口结构体
type RoleUpdateApisReq struct {
	RoleID uint   `json:"roleId" validate:"required"`
	ApiIds []uint `json:"apiIds" validate:"required"`
}
