package middleware

import (
	"GinBlog/utils"
	"GinBlog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// 生成token
func (this *MyClaims) SetToken(username string) (string, int)  {
	expireTime := time.Now().Add(10 *time.Hour) //有效时间
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //unix() 生成int64时间戳
			Issuer: "ginblog", //签发人
		},
	}
	requestClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims) //签发token。SigningMethodHS256是签发的方法
	token, err := requestClaim.SignedString(JwtKey)
	if err != nil {
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS
}

// 验证token
func (this *MyClaims) CheckToken(token string) (*MyClaims, int) {
	settoken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})//生成token
	if key, _ := settoken.Claims.(*MyClaims); settoken.Valid {
		return key, errmsg.SUCCESS
	}else {
		return nil, errmsg.ERROR
	}//比对token
}


// jwt中间件
func JwtToken() gin.HandlerFunc {
	var this MyClaims
	return func(c *gin.Context) {
		tokenHerder := c.Request.Header.Get("Authorization")
		code := errmsg.SUCCESS
		if tokenHerder == "" {
			code = errmsg.ErrorTokenExist
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHerder," ",2)
		if len(checkToken) !=2 &&checkToken[0] !="Bearer" {
			code = errmsg.ErrorToeknTypeWrong
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		} // 判断token格式是否错误
		key, tCode := this.CheckToken(checkToken[1])
		if tCode == errmsg.ERROR {
			code = errmsg.ErrorTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}// 判断token是否错误
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.JSON(http.StatusOK, gin.H{
				"code":code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}// 判断token是否过期
		c.Set("username", key.Username)
		c.Next()
	}
}