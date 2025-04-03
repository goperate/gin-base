package main

import (
	"app/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.SetupUserRoutes(r)
	err := r.Run(":8080")
	if err != nil {
		panic("gin启动失败, " + err.Error())
	}
}
