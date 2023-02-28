package controller

import (
	"Glog/dao"
	"Glog/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(301, "/")
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.Mgr.Login(username)

	if user.Username == "" {
		c.HTML(200, "login.html", "用户名不存在")
	} else {
		if user.Password != password {
			c.HTML(200, "login.html", "密码错误")
		} else {
			//c.String(200, "登录成功", nil)
			fmt.Println("登录成功")
			c.Redirect(301, "/")
		}
	}
}
