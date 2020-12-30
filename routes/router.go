package routes

import (
	v1 "GinBlog/api/v1"
	"GinBlog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default() // Default加了日志中间件

	router := r.Group("api/v1")
	{
		//v1.GET("hello", func(context *gin.Context) {
		//	context.JSON(http.StatusOK, gin.H{
		//		"msg": "ok",
		//	})
		//})

		//用户模块的路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeletUser)
		//分类模块的路由接口
		router.POST("category/add", v1.AddCategory)

		//文章模块的路由接口

	}

	_ = r.Run(utils.HttpPort)
}