package routes

import (
	"github.com/eryajf/go-ldap-admin/controller"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("ping", controller.Demo)
		base.GET("encryptpwd", controller.Base.EncryptPasswd) // 生成加密密码
		base.GET("decryptpwd", controller.Base.DecryptPasswd) // 密码解密为明文
		// 登录登出刷新token无需鉴权
		base.POST("/login", authMiddleware.LoginHandler)
		base.POST("/logout", authMiddleware.LogoutHandler)
		base.POST("/refreshToken", authMiddleware.RefreshHandler)
		base.POST("/sendcode", controller.Base.SendCode)   // 给用户邮箱发送验证码
		base.POST("/changePwd", controller.Base.ChangePwd) // 修改用户密码
		base.GET("/dashboard", controller.Base.Dashboard)  // 系统首页展示数据
	}
	return r
}
