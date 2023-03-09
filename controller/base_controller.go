package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// SendCode 给用户邮箱发送验证码
func (m *BaseController) SendCode(c *gin.Context) {
	req := new(request.BaseSendCodeReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.SendCode(c, req)
	})
}

// ChangePwd 用户通过邮箱修改密码
func (m *BaseController) ChangePwd(c *gin.Context) {
	req := new(request.BaseChangePwdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.ChangePwd(c, req)
	})
}

// Dashboard 系统首页展示数据
func (m *BaseController) Dashboard(c *gin.Context) {
	req := new(request.BaseDashboardReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.Dashboard(c, req)
	})
}

// EncryptPasswd 生成加密密码
func (m *BaseController) EncryptPasswd(c *gin.Context) {
	req := new(request.EncryptPasswdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.EncryptPasswd(c, req)
	})
}

// DecryptPasswd 密码解密为明文
func (m *BaseController) DecryptPasswd(c *gin.Context) {
	req := new(request.DecryptPasswdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.DecryptPasswd(c, req)
	})
}
