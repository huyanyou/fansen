package server

import (
	"fansenbk/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()
	// 创建api分组
	api := r.Group("/api")
	admin := api.Group("/admin")
	admin.POST("/", controller.CreateAdmin)
	admin.PUT("/:id", controller.UpdateAdminById)
	// 根据Id获取管理员
	admin.GET("/:id", controller.GetAdminById)
	admin.DELETE("/:id", controller.DeleteAdminById)

	customer := api.Group("/customer")
	customer.POST("/", controller.CreateCustomer)
	r.Run()
}
