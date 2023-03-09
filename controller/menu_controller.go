package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// GetTree 菜单树
func (m *MenuController) GetTree(c *gin.Context) {
	req := new(request.MenuGetTreeReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.GetTree(c, req)
	})
}

// GetUserMenuTreeByUserId 获取用户菜单树
func (m *MenuController) GetAccessTree(c *gin.Context) {
	req := new(request.MenuGetAccessTreeReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.GetAccessTree(c, req)
	})
}

// Add 新建
func (m *MenuController) Add(c *gin.Context) {
	req := new(request.MenuAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.Add(c, req)
	})
}

// Update 更新记录
func (m *MenuController) Update(c *gin.Context) {
	req := new(request.MenuUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.Update(c, req)
	})
}

// Delete 删除记录
func (m *MenuController) Delete(c *gin.Context) {
	req := new(request.MenuDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.Delete(c, req)
	})
}
