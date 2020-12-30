package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

//添加分类
func AddCategory(c *gin.Context)  {
	var data model.Category
	_ = c.ShouldBind(&data)
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		model.AddCategory(&data)
	}
	if code == errmsg.ErrorCategoryUsed {
		code = errmsg.ErrorCategoryUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类

//查询单个分类下的文章

//查询分类列表

//编辑分类

//删除分类