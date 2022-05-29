package routes

import (
	"github.com/eryajf/go-ldap-admin/controller"
	"github.com/eryajf/go-ldap-admin/middleware"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitApiRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	api := r.Group("/api")
	// 开启jwt认证中间件
	api.Use(authMiddleware.MiddlewareFunc())
	// 开启casbin鉴权中间件
	api.Use(middleware.CasbinMiddleware())
	{
		api.GET("/tree", controller.Api.GetTree)
		api.GET("/list", controller.Api.List)
		api.POST("/add", controller.Api.Add)
		api.POST("/update", controller.Api.Update)
		api.POST("/delete", controller.Api.Delete)
	}

	return r
}
