package router

import (
	"Glog/controller"
	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()

	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.GET("/login", controller.GoLogin)
	e.POST("login", controller.Login)
	e.GET("/", controller.Index)

	e.Run(":8080")
}
