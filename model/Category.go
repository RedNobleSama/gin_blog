package model

import (
	"GinBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 分类是否存在
func CheckCategory(name string) int {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.SUCCESS
}

// 创建分类
func AddCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表

// 查询分类下的文章

// 编辑分类

// 删除分类