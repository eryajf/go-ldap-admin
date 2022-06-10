package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

// List 记录列表
func (m *ApiController) List(c *gin.Context) {
	req := new(request.ApiListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Api.List(c, req)
	})
}

// GetTree 接口树
func (m *ApiController) GetTree(c *gin.Context) {
	req := new(request.ApiGetTreeReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Api.GetTree(c, req)
	})
}

// Add 新建记录
func (m *ApiController) Add(c *gin.Context) {
	req := new(request.ApiAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Api.Add(c, req)
	})
}

// Update 更新记录
func (m *ApiController) Update(c *gin.Context) {
	req := new(request.ApiUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Api.Update(c, req)
	})
}

// Delete 删除记录
func (m *ApiController) Delete(c *gin.Context) {
	req := new(request.ApiDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Api.Delete(c, req)
	})
}
