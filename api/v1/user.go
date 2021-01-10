package v1

import (
	"GinBlog/model"
	validator2 "GinBlog/utils/validator"
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



//添加用户
func AddUser(c *gin.Context) {
	fmt.Println("添加用户")
	var user model.User
	var msg string
	var code int
	_ = c.ShouldBind(&user)

	msg, code = validator2.Validate(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"message": msg,
		})
		c.Abort()
	}

	code = user.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		user.CreateUser(&user)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status" : code,
		"data": user,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询用户列表
func GetUsers(c *gin.Context) {
	fmt.Println("查询用户")
	var user model.User
	//strconv.Atoi 将String转为Int
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum ==0 {
		pageNum = -1
	}
	data := user.GetUsers(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"msg": errmsg.GetErrMsg(code),
	})
}

//编辑用户
func EditUser(c *gin.Context) {
	fmt.Println("编辑用户")
	var user model.User
	_ = c.ShouldBind(&user)
	code := user.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		id, _ := strconv.Atoi(c.Param("id"))
		fmt.Println(user)
		user.EditUser(id, &user)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data" : user,
		"msg": errmsg.GetErrMsg(code),
	})
}

//删除用户
func DeletUser(c *gin.Context) {
	fmt.Println("删除用户")
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	code := user.DeleteUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg": errmsg.GetErrMsg(code),
	})
}
