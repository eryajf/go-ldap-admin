package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type RoleController struct{}

// List 记录列表
func (m *RoleController) List(c *gin.Context) {
	req := new(request.RoleListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.List(c, req)
	})
}

// Add 新建
func (m *RoleController) Add(c *gin.Context) {
	req := new(request.RoleAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.Add(c, req)
	})
}

// Update 更新记录
func (m *RoleController) Update(c *gin.Context) {
	req := new(request.RoleUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.Update(c, req)
	})
}

// Delete 删除记录
func (m *RoleController) Delete(c *gin.Context) {
	req := new(request.RoleDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.Delete(c, req)
	})
}

// GetMenuList 获取菜单列表
func (m *RoleController) GetMenuList(c *gin.Context) {
	req := new(request.RoleGetMenuListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.GetMenuList(c, req)
	})
}

// GetApiList 获取接口列表
func (m *RoleController) GetApiList(c *gin.Context) {
	req := new(request.RoleGetApiListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.GetApiList(c, req)
	})
}

// UpdateMenus 更新菜单
func (m *RoleController) UpdateMenus(c *gin.Context) {
	req := new(request.RoleUpdateMenusReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.UpdateMenus(c, req)
	})
}

// UpdateApis 更新接口
func (m *RoleController) UpdateApis(c *gin.Context) {
	req := new(request.RoleUpdateApisReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Role.UpdateApis(c, req)
	})
}
