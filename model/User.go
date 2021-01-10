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
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required, min=4, max=12" label:"用户名"`
	Password string `gorm:"type:varchar(50);not null" json:"password" validate:"required, min=6, max=20" label:"密码"`
	Role     int    `gorm:"type:int;Default: 2" json:"role" validate:"required, gte=2" label:"角色"`
}

// 查询用户是否存在
func (u User) CheckUser(username string) int {
	db.Select("id").Where("username = ?", username).First(&u)
	if u.ID > 0 {
		return errmsg.ErrorUsernameUsed
	}
	fmt.Println(u.ID)
	return errmsg.SUCCESS
}

// 新增用户
func (u User) CreateUser(user *User) int {
	err := db.Create(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询用户列表
// 返回User模型的切片
func (u User) GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}

// 编辑用户信息
func (u User) EditUser(id int, user *User) int {
	err := db.Model(&user).Where("id=?", id).Updates(map[string]interface{}{"username": user.Username, "role": user.Role})
	if err.Error != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func (u User) DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 钩子函数密码加密
func (u *User) BeforeSave() {
	u.Password = string(HashPassword([]byte(u.Password)))
}

// 用户密码加密
func HashPassword(password []byte) []byte {
	hashPWS, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		fmt.Println("密码格式有问题", err.Error())
	}
	return hashPWS
}

// 密码校验
func ComparePassword(hashedPassword string, password string) bool {
	byteHash := []byte(hashedPassword)
	bytePassword := []byte(password)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePassword)
	if err != nil {
		fmt.Println("密码错误", err.Error())
		return false
	}
	return true
}

// 登录
func (u *User)Login(username string, password string) int {
	var user User
	num := db.Where("username = ?", username).Find(&user).RowsAffected
	fmt.Println(num)
	if num == 0 {
		return errmsg.ErrorUserNotExist
	}

	check := ComparePassword(user.Password, password)
	if !check {
		return errmsg.ErrorPasswordWrong
	} else {
		return errmsg.SUCCESS
	}
}