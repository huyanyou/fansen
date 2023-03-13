package server

import (
	"fansenbk/controller"
	"fansenbk/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	// 解决redirecting request 307
	r.Use(middleware.Cors())
	// 创建api分组
	api := r.Group("/api")
	admin := api.Group("/admin")
	admin.POST("", controller.CreateAdmin)
	admin.PUT("/:id", controller.UpdateAdminById)
	// 根据Id获取管理员
	admin.GET("/:id", controller.GetAdminById)
	admin.DELETE("/:id", controller.DeleteAdminById)

	customer := api.Group("/customer")
	customer.POST("", controller.CreateCustomer)
	customer.POST("/dels", controller.DelCustomeres)
	customer.POST("updates", controller.UpdateCustomeres)
	// 允许代理跨域
	r.SetTrustedProxies(nil)
	r.Run(":9000")
}
