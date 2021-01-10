package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var article model.Article
//添加文章
// 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := data.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}



//查询单个文章
func GetOneArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	data := article.GetOneArticle(id)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章列表
func GetArticles(c *gin.Context) {
	var article model.Article
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code := article.GetArticles(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}


//编辑文章
func EditArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBind(&article)
	code := article.EditArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeletArticle(c *gin.Context) {
	var article model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	code := article.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": errmsg.GetErrMsg(code),
	})
}
