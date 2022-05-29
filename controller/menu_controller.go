package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/svc/request"

	"github.com/gin-gonic/gin"
)

type MenuController struct{}

// // List 记录列表
// func (m *MenuController) List(c *gin.Context) {
// 	req := new(request.MenuListReq)
// 	Run(c, req, func() (interface{}, interface{}) {
// 		return logic.Menu.List(c, req)
// 	})
// }

// GetTree 菜单树
func (m *MenuController) GetTree(c *gin.Context) {
	req := new(request.MenuGetTreeReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Menu.GetTree(c, req)
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
