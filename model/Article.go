package model

import (
	"GinBlog/utils/errmsg"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title    string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid      int      `gorm:"type:int;not null" json:"cid"`
	Desc     string   `gorm:"type:varchar(200)" json:"desc"`
	Content  string   `gorm:"type:longtext" json:"content"`
	Image    string   `gorm:"type:varchar(100)" json:"image"`
}

//添加文章
func (this Article) CreateArt(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

//查询单篇文章
func (this Article)GetOneArticle (id int) Article {
	var article Article
	err := db.Preload("Category").Where("id=?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article
	}
	return article
}

//查询文章列表
func (this Article)GetArticles (pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errmsg.ERROR
	}
	return articles, errmsg.SUCCESS
}


// 编辑文章
func (this *Article)EditArticle(id int, article *Article) int {
	err := db.Model(&this).Where("id=?", id).Updates(map[string]interface{}{"title": article.Title, "desc": article.Desc, "content": article.Content}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func (this Article) DeleteArticle(id int) int  {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
