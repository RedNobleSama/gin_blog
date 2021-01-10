package model

import (
	"GinBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 分类是否存在
func (c Category) CheckCategory(name string) int {
	db.Select("id").Where("name = ?", name).First(&c)
	if c.ID > 0 {
		return errmsg.ErrorCategoryUsed
	}
	return errmsg.ErrorCategoryNotExist
}

// 创建分类
func (c Category) AddCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 查询分类列表
func (c Category) GetCategorys(pageSize int, pageNum int) []Category {
	var category []Category
	err := db.Limit(pageSize).Offset((pageNum-1) * pageSize).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil
	}
	return category
}

// 查询分类下的文章
func (category *Category) GetCategoryArt(id int, pageSize int, pageNum int) ([]Article, int) {
	var article []Article
	err := db.Where("id = ?", id).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	db.Model(&article).Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&article)
	return article, errmsg.SUCCESS
}


// 编辑分类
func (c Category) EditCategory(id int, data *Category) int {
	var category Category
	err := db.Model(&category).Where("id=?", id).Updates(map[string]interface{}{"name": data.Name}).Error
	if err != nil{
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}



// 删除分类

func (c Category) DeleteCategory(id int) int {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}