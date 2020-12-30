package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加用户
func AddUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个用户

//查询用户列表
func GetUsers(c *gin.Context) {
	//strconv.Atoi 将String转为Int
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum ==0 {
		pageNum = -1
	}
	data := model.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	_ = c.ShouldBind(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		id, _ := strconv.Atoi(c.Param("id"))
		model.EditUser(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data" : data,
		"msg": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeletUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": errmsg.GetErrMsg(code),
	})
}