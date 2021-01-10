package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger()) // 引入重写后的日志中间件
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) //加载中间件
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeletUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeletCategory)

		//文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeletArticle)

		// 文件上传
		auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.POST("login", v1.UserLogin)
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("categorys", v1.GetCategorys)
		router.GET("category/:id", v1.GetCategoryArt)
		router.GET("articles", v1.GetArticles)
		router.GET("article/:id", v1.GetOneArticle)

	}

	//启动服务
	return r
}