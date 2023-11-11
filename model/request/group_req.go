package request

// GroupListReq 获取资源列表结构体
type GroupListReq struct {
	GroupName string `json:"groupName" form:"groupName"`
	Remark    string `json:"remark" form:"remark"`
	PageNum   int    `json:"pageNum" form:"pageNum"`
	PageSize  int    `json:"pageSize" form:"pageSize"`
	SyncState uint   `json:"syncState" form:"syncState"`
}

// GroupListAllReq 获取资源列表结构体，不分页
type GroupListAllReq struct {
	GroupName          string `json:"groupName" form:"groupName"`
	GroupType          string `json:"groupType" form:"groupType"`
	Remark             string `json:"remark" form:"remark"`
	Source             string `json:"source" form:"source"`
	SourceDeptId       string `json:"sourceDeptId"`
	SourceDeptParentId string `json:"SourceDeptParentId"`
}

// GroupAddReq 添加资源结构体
type GroupAddReq struct {
	GroupType string `json:"groupType" validate:"required,min=1,max=20"`
	GroupName string `json:"groupName" validate:"required,min=1,max=128"`
	//父级Id 大于等于0 必填
	ParentId uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark   string `json:"remark" validate:"min=0,max=128"` // 分组的中文描述
}

// DingTalkGroupAddReq 添加钉钉资源结构体
type DingGroupAddReq struct {
	GroupType string `json:"groupType" validate:"required,min=1,max=20"`
	GroupName string `json:"groupName" validate:"required,min=1,max=128"`
	//父级Id 大于等于0 必填
	ParentId           uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark             string `json:"remark" validate:"min=0,max=128"` // 分组的中文描述
	SourceDeptId       string `json:"sourceDeptId"`
	Source             string `json:"source"`
	SourceDeptParentId string `json:"SourceDeptParentId"`
	SourceUserNum      int    `json:"sourceUserNum"`
}

// WeComGroupAddReq 添加企业微信资源结构体
type WeComGroupAddReq struct {
	GroupType string `json:"groupType" validate:"required,min=1,max=20"`
	GroupName string `json:"groupName" validate:"required,min=1,max=128"`
	//父级Id 大于等于0 必填
	ParentId           uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark             string `json:"remark" validate:"min=0,max=128"` // 分组的中文描述
	SourceDeptId       string `json:"sourceDeptId"`
	Source             string `json:"source"`
	SourceDeptParentId string `json:"SourceDeptParentId"`
	SourceUserNum      int    `json:"sourceUserNum"`
}

// GroupUpdateReq 更新资源结构体
type GroupUpdateReq struct {
	ID        uint   `json:"id" form:"id" validate:"required"`
	GroupName string `json:"groupName" validate:"required,min=1,max=128"`
	Remark    string `json:"remark" validate:"min=0,max=128"` // 分组的中文描述
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

// SyncDingTalkDeptsReq 同步钉钉部门信息
type SyncDingTalkDeptsReq struct {
}

// SyncWeComDeptsReq 同步企业微信部门信息
type SyncWeComDeptsReq struct {
}

// SyncFeiShuDeptsReq 同步飞书部门信息
type SyncFeiShuDeptsReq struct {
}

// SyncOpenLdapDeptsReq 同步原ldap部门信息
type SyncOpenLdapDeptsReq struct {
}

// SyncOpenLdapDeptsReq 同步原ldap部门信息
type SyncSqlGrooupsReq struct {
	GroupIds []uint `json:"groupIds" validate:"required"`
}
