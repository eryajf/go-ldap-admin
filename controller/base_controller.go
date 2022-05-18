package controller

import (
	"github.com/eryajf-world/go-ldap-admin/logic"
	"github.com/eryajf-world/go-ldap-admin/svc/request"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

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
