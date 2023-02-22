package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

// Add 添加记录
func (m *UserController) Add(c *gin.Context) {
	req := new(request.UserAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.Add(c, req)
	})
}

// Update 更新记录
func (m *UserController) Update(c *gin.Context) {
	req := new(request.UserUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.Update(c, req)
	})
}

// List 记录列表
func (m *UserController) List(c *gin.Context) {
	req := new(request.UserListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.List(c, req)
	})
}

// Delete 删除记录
func (m UserController) Delete(c *gin.Context) {
	req := new(request.UserDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.Delete(c, req)
	})
}

// ChangePwd 更新密码
func (m UserController) ChangePwd(c *gin.Context) {
	req := new(request.UserChangePwdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.ChangePwd(c, req)
	})
}

// ChangeUserStatus 更改用户状态
func (m UserController) ChangeUserStatus(c *gin.Context) {
	req := new(request.UserChangeUserStatusReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.ChangeUserStatus(c, req)
	})
}

// GetUserInfo 获取当前登录用户信息
func (uc UserController) GetUserInfo(c *gin.Context) {
	req := new(request.UserGetUserInfoReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.User.GetUserInfo(c, req)
	})
}

// 同步钉钉用户信息
func (uc UserController) SyncDingTalkUsers(c *gin.Context) {
	req := new(request.SyncDingUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.DingTalk.SyncDingTalkUsers(c, req)
	})
}

// 同步企业微信用户信息
func (uc UserController) SyncWeComUsers(c *gin.Context) {
	req := new(request.SyncWeComUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.WeCom.SyncWeComUsers(c, req)
	})
}

// 同步飞书用户信息
func (uc UserController) SyncFeiShuUsers(c *gin.Context) {
	req := new(request.SyncFeiShuUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FeiShu.SyncFeiShuUsers(c, req)
	})
}

// 同步ldap用户信息
func (uc UserController) SyncOpenLdapUsers(c *gin.Context) {
	req := new(request.SyncOpenLdapUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.OpenLdap.SyncOpenLdapUsers(c, req)
	})
}

// 同步sql用户信息到ldap
func (uc UserController) SyncSqlUsers(c *gin.Context) {
	req := new(request.SyncSqlUserReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Sql.SyncSqlUsers(c, req)
	})
}
