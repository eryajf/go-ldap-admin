package request

// UserAddReq 创建资源结构体
type UserAddReq struct {
	Username      string `json:"username" validate:"required,min=2,max=20"`
	Password      string `json:"password"`
	Nickname      string `json:"nickname" validate:"required,min=0,max=20"`
	GivenName     string `json:"givenName" validate:"min=0,max=20"`
	Mail          string `json:"mail" validate:"required,min=0,max=20"`
	JobNumber     string `json:"jobNumber" validate:"required,min=0,max=20"`
	PostalAddress string `json:"postalAddress" validate:"min=0,max=255"`
	Departments   string `json:"departments" validate:"min=0,max=255"`
	Position      string `json:"position" validate:"min=0,max=255"`
	Mobile        string `json:"mobile" validate:"required,checkMobile"`
	Avatar        string `json:"avatar"`
	Introduction  string `json:"introduction" validate:"min=0,max=255"`
	Status        uint   `json:"status" validate:"oneof=1 2"`
	RoleIds       []uint `json:"roleIds" validate:"required"`
}

// UserUpdateReq 更新资源结构体
type UserUpdateReq struct {
	ID            uint   `json:"id" validate:"required"`
	Nickname      string `json:"nickname" validate:"min=0,max=20"`
	GivenName     string `json:"givenName" validate:"min=0,max=20"`
	Mail          string `json:"mail" validate:"min=0,max=20"`
	JobNumber     string `json:"jobNumber" validate:"min=0,max=20"`
	PostalAddress string `json:"postalAddress" validate:"min=0,max=255"`
	Departments   string `json:"departments" validate:"min=0,max=255"`
	Position      string `json:"position" validate:"min=0,max=255"`
	Mobile        string `json:"mobile" validate:"checkMobile"`
	Avatar        string `json:"avatar"`
	Introduction  string `json:"introduction" validate:"min=0,max=255"`
	RoleIds       []uint `json:"roleIds" validate:"required"`
}

// UserDeleteReq 批量删除资源结构体
type UserDeleteReq struct {
	UserIds []uint `json:"userIds" validate:"required"`
}

// UserChangePwdReq 修改密码结构体
type UserChangePwdReq struct {
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}

// UserChangeUserStatusReq 修改用户状态结构体
type UserChangeUserStatusReq struct {
	ID     uint `json:"id" validate:"required"`
	Status uint `json:"status" validate:"oneof=1 2"`
}

// UserGetUserInfoReq 获取用户信息结构体
type UserGetUserInfoReq struct {
}

// UserListReq 获取用户列表结构体
type UserListReq struct {
	Username string `json:"username" form:"username"`
	Mobile   string `json:"mobile" form:"mobile" `
	Nickname string `json:"nickname" form:"nickname"`
	Status   uint   `json:"status" form:"status" `
	PageNum  int    `json:"pageNum" form:"pageNum"`
	PageSize int    `json:"pageSize" form:"pageSize"`
}

// RegisterAndLoginReq 用户登录结构体
type RegisterAndLoginReq struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
