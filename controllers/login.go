package controllers

import (
	"fmt"
	"gin_blog/logger"
	"gin_blog/models"
	"gin_blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"go.uber.org/zap"
	"net/http"
)

func LoginGet(c *gin.Context) {
	//返回html
	c.HTML(http.StatusOK, "login.html", gin.H{"title": "登录页"})
}

func LoginPost(c *gin.Context){
	// 取出请求数据
	// 校验用户名密码是否正确
	// 返回响应
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, ",password:", password)
	logger.Debug("login",zap.String("username", username),zap.String("password", password))

	id := models.QueryUserWithParam(username, utils.MD5(password)) // 去数据库查
	fmt.Println("id:", id)
	if id > 0 {
		// 给响应种上Cookie
		session := sessions.Default(c)
		session.Set("login_user", username)  // 在session中保存k-v,然后写入cookie
		session.Save()

		//c.Redirect(http.StatusFound, "/index") // 浏览器收到这个就会跳转到我指定的页面
		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}


func LogoutHandler(c *gin.Context){
	//清除该用户登录状态的数据
	session := sessions.Default(c)
	session.Delete("login_user")
	session.Save()

	c.Redirect(http.StatusFound, "/login")
}