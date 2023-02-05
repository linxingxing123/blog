package routers

import (
	"gin_blog/controllers"
	"gin_blog/logger"
	"gin_blog/middlewares"
	"github.com/gin-contrib/sessions" // session包 定义了一套session操作的接口 类似于 database/sql
	"github.com/gin-gonic/gin"
	"html/template"
	"time"

	//"github.com/gin-contrib/sessions/cookie"  // session具体存储的介质
	"github.com/gin-contrib/sessions/redis" // session具体存储的介质
	//"github.com/gin-contrib/sessions/memcached"  // session具体存储的介质

	// github.com/go-redis/redis  --> go连接redis的一个第三方库
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))

	//时间转换
	r.SetFuncMap(template.FuncMap{
		"timeStr": func(timestamp int64) string {
			return time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
		},
	})

	r.Static("/static", "static")
	r.LoadHTMLGlob("views/*")

	// 设置session midddleware
	//store := cookie.NewStore([]byte("secret"))
	store, _ := redis.NewStore(10, "tcp", "127.0.0.1:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// 登录注册 无需认证
	{
		r.GET("/register", controllers.RegisterGet)
		r.POST("/register", controllers.RegisterPost)

		r.GET("/login", controllers.LoginGet)
		r.POST("/login", controllers.LoginPost)

		// topN
		r.GET("/article/top/:n", controllers.ArticleTopN)
	}

	{
		basicAuthGroup := r.Group("/", middlewares.BasicAuth()) // 路由组注册中间件
		basicAuthGroup.GET("/home", controllers.HomeGet)
		basicAuthGroup.GET("/", controllers.IndexHandler)
		basicAuthGroup.GET("/logout", controllers.LogoutHandler)

		//路由组
		article := basicAuthGroup.Group("/article")
		{
			article.GET("/add", controllers.AddArticleGet)
			article.POST("/add", controllers.AddArticlePost)
			// 文章详情
			article.GET("/show/:id", controllers.ShowArticleGet)
			// 更新文章
			article.GET("/update", controllers.UpdateArticleGet)
			article.POST("/update", controllers.UpdateArticlePost)
			// 删除文章
			article.GET("/delete", controllers.DeleteArticle)

		}
		// 相册
		basicAuthGroup.GET("/album", controllers.AlbumGet)
		// 文件上传
		basicAuthGroup.POST("/upload", controllers.UploadPost)

	}

	return r

}