package model

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"primary_key auto_increment not null unique"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var DB *gorm.DB

// 创建连接数据库
func InitDB() {
	db, err := gorm.Open(sqlite.Open("fansenbk.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Transaction{}, &Admin{}, &Customer{})
	DB = db
}
