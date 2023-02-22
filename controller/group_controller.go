package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type GroupController struct{}

// List 记录列表
func (m *GroupController) List(c *gin.Context) {
	req := new(request.GroupListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.List(c, req)
	})
}

// UserInGroup 在分组内的用户
func (m *GroupController) UserInGroup(c *gin.Context) {
	req := new(request.UserInGroupReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.UserInGroup(c, req)
	})
}

// UserNoInGroup 不在分组的用户
func (m *GroupController) UserNoInGroup(c *gin.Context) {
	req := new(request.UserNoInGroupReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.UserNoInGroup(c, req)
	})
}

// GetTree 接口树
func (m *GroupController) GetTree(c *gin.Context) {
	req := new(request.GroupListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.GetTree(c, req)
	})
}

// Add 新建记录
func (m *GroupController) Add(c *gin.Context) {
	req := new(request.GroupAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.Add(c, req)
	})
}

// Update 更新记录
func (m *GroupController) Update(c *gin.Context) {
	req := new(request.GroupUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.Update(c, req)
	})
}

// Delete 删除记录
func (m *GroupController) Delete(c *gin.Context) {
	req := new(request.GroupDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.Delete(c, req)
	})
}

// AddUser 添加用户
func (m *GroupController) AddUser(c *gin.Context) {
	req := new(request.GroupAddUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.AddUser(c, req)
	})
}

// RemoveUser 移除用户
func (m *GroupController) RemoveUser(c *gin.Context) {
	req := new(request.GroupRemoveUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Group.RemoveUser(c, req)
	})
}

//同步钉钉部门信息
func (m *GroupController) SyncDingTalkDepts(c *gin.Context) {
	req := new(request.SyncDingTalkDeptsReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.DingTalk.SyncDingTalkDepts(c, req)
	})
}

//同步企业微信部门信息
func (m *GroupController) SyncWeComDepts(c *gin.Context) {
	req := new(request.SyncWeComDeptsReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.WeCom.SyncWeComDepts(c, req)
	})
}

//同步飞书部门信息
func (m *GroupController) SyncFeiShuDepts(c *gin.Context) {
	req := new(request.SyncFeiShuDeptsReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FeiShu.SyncFeiShuDepts(c, req)
	})
}

//同步原ldap部门信息
func (m *GroupController) SyncOpenLdapDepts(c *gin.Context) {
	req := new(request.SyncOpenLdapDeptsReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.OpenLdap.SyncOpenLdapDepts(c, req)
	})
}

//同步Sql中的分组信息到ldap
func (m *GroupController) SyncSqlGroups(c *gin.Context) {
	req := new(request.SyncSqlGrooupsReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Sql.SyncSqlGroups(c, req)
	})
}
