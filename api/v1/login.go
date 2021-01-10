package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户登录
func UserLogin(c *gin.Context) {
	var user model.User
	var myclaims middleware.MyClaims
	_ = c.ShouldBind(&user)
	var token string
	var code int
	code = user.Login(user.Username, user.Password)
	if code == errmsg.SUCCESS{
		token, code = myclaims.SetToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":code,
		"message": errmsg.GetErrMsg(code),
		"token": token,
	})

}
