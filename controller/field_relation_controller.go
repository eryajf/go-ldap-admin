package controller

import (
	"github.com/eryajf/go-ldap-admin/logic"
	"github.com/eryajf/go-ldap-admin/model/request"

	"github.com/gin-gonic/gin"
)

type FieldRelationController struct{}

// List 记录列表
func (m *FieldRelationController) List(c *gin.Context) {
	req := new(request.FieldRelationListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FieldRelation.List(c, req)
	})
}

// Add 新建记录
func (m *FieldRelationController) Add(c *gin.Context) {
	req := new(request.FieldRelationAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FieldRelation.Add(c, req)
	})
}

// Update 更新记录
func (m *FieldRelationController) Update(c *gin.Context) {
	req := new(request.FieldRelationUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FieldRelation.Update(c, req)
	})
}

// Delete 删除记录
func (m *FieldRelationController) Delete(c *gin.Context) {
	req := new(request.FieldRelationDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.FieldRelation.Delete(c, req)
	})
}
