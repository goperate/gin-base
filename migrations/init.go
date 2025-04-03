package main

import (
	"app/config"
	"app/models"
)

func main() {
	config.ConnectDB()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("数据库初始化失败, " + err.Error())
	}
	println("数据库初始化成功")
}
