package request

// GroupListReq 获取资源列表结构体
type GroupListReq struct {
	GroupName string `json:"groupName" form:"groupName"`
	Remark    string `json:"remark" form:"remark"`
	PageNum   int    `json:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

// GroupAddReq 添加资源结构体
type GroupAddReq struct {
	GroupType string `json:"groupType" validate:"required,min=1,max=20"`
	GroupName string `json:"groupName" validate:"required,min=1,max=20"`
	//父级Id 大于等于0 必填
	ParentId uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark   string `json:"remark" validate:"min=0,max=100"` // 分组的中文描述
}

// GroupUpdateReq 更新资源结构体
type GroupUpdateReq struct {
	ID        uint   `json:"id" form:"id" validate:"required"`
	GroupName string `json:"groupName" validate:"required,min=1,max=20"`
	Remark    string `json:"remark" validate:"min=0,max=100"` // 分组的中文描述
}

// GroupDeleteReq 删除资源结构体
type GroupDeleteReq struct {
	GroupIds []uint `json:"groupIds" validate:"required"`
}

// GroupGetTreeReq 获取资源树结构体
type GroupGetTreeReq struct {
	GroupName string `json:"groupName" form:"groupName"`
	Remark    string `json:"remark" form:"remark"`
	PageNum   int    `json:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
}

type GroupAddUserReq struct {
	GroupID uint   `json:"groupId" validate:"required"`
	UserIds []uint `json:"userIds" validate:"required"`
}

type GroupRemoveUserReq struct {
	GroupID uint   `json:"groupId" validate:"required"`
	UserIds []uint `json:"userIds" validate:"required"`
}

// UserInGroupReq 在分组内的用户
type UserInGroupReq struct {
	GroupID  uint   `json:"groupId" form:"groupId" validate:"required"`
	Nickname string `json:"nickname" form:"nickname"`
}

// UserNoInGroupReq 不在分组内的用户
type UserNoInGroupReq struct {
	GroupID  uint   `json:"groupId" form:"groupId" validate:"required"`
	Nickname string `json:"nickname" form:"nickname"`
}
