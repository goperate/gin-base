package main

import (
	"app/config"
	"app/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupUserRoutes(r)
	config.CONF.SetDefault("port", "80")
	err := r.Run(":" + config.CONF.GetString("port"))
	if err != nil {
		panic("gin启动失败, " + err.Error())
	}
}
