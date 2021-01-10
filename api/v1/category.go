package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context)  {
	var category model.Category
	_ = c.ShouldBind(&category)
	code := category.CheckCategory(category.Name)
	if code == errmsg.ErrorCategoryNotExist {
		category.AddCategory(&category)
	}
	if code == errmsg.ErrorCategoryUsed {
		code = errmsg.ErrorCategoryUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": category,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个分类

//查询单个分类下的文章

//查询分类列表
func GetCategorys(c *gin.Context) {
	var category model.Category
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data := category.GetCategorys(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})

}

func GetCategoryArt(c *gin.Context){
	var category model.Category
	id,_ := strconv.Atoi(c.Param("id"))
	fmt.Println(id)
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code := category.GetCategoryArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {
	var category model.Category
	id,_ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&category)
	code := category.CheckCategory(category.Name)
	if code == errmsg.ErrorCategoryNotExist{
		category.EditCategory(id, &category)
	}
	if code == errmsg.ErrorCategoryUsed {
		code =  errmsg.ErrorCategoryUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": category,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeletCategory(c *gin.Context) {
	var category model.Category
	id,_ := strconv.Atoi(c.Param("id"))
	code := category.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"message": errmsg.GetErrMsg(code),
	})
}