package main

import (
	"fansenbk/model"
	"fansenbk/server"
)

func init() {
	model.InitDB()
}
func main() {
	server.NewRouter()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	// db, err := gorm.Open(sqlite.Open("fansenbk.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// db.AutoMigrate(&model.Transaction{}, &model.Admin{}, &model.Customer{})
}
