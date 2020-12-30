package model

import (
	"GinBlog/utils/errmsg"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	//GORM 定义一个 gorm.Model 结构体，其包括字段 ID、CreatedAt、UpdatedAt、DeletedAt
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(50);not null" json:"password"`
	Role int `gorm:"type:int" json:"role"`
}

// 查询用户是否存在
func CheckUser(username string) int {
	var users User
	db.Select("id").Where("username = ?", username).First(&users)
	if users.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) int {
	//data.Password = string(HashPassword([]byte(data.Password)))
	err := db.Create(&data).Error
	if err !=nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
// 返回User模型的切片
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum-1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return users
}

// 钩子函数密码加密
func (u *User)BeforeSave() {
	u.Password = string(HashPassword([]byte(u.Password)))
}


// 用户密码加密
func HashPassword(password []byte) []byte {
	hashPWS, err := bcrypt.GenerateFromPassword(password,10)
	if err != nil{
		fmt.Println("密码格式有问题", err.Error())
	}
	return hashPWS
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	err := db.Model(&user).Where("id=?", id).Updates(map[string]interface{}{"username": data.Username, "role": data.Role})
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}


// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}