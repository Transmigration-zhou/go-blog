package model

import (
	"go-blog/utils/errormessage"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignKey:Cid"`
	Title       string   `gorm:"type:varchar(20);not null" json:"title"`
	Cid         uint     `gorm:"type:uint;not null;" json:"cid"`
	Description string   `gorm:"type:varchar(100);not null" json:"description"`
	Content     string   `gorm:"type:longtext" json:"content"`
	Image       string   `gorm:"type:varchar(100)" json:"image"`
}

// CreateArticle 新建文章
func CreateArticle(article *Article) int {
	err := db.Create(&article).Error
	if err != nil {
		return errormessage.ERROR
		log.Fatal(err)
	}
	return errormessage.SUCCESS
}

// DeleteArticle 删除指定id的用户
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}

// GetCategoryArticle 查询某分类cid下所有文章
func GetCategoryArticle(id int, pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&articles).Error
	if err != nil {
		return nil, errormessage.ERROR_CATE_NOT_EXIST
	}
	return articles, errormessage.SUCCESS
}

// GetSingleArticle 查询单个文章
func GetSingleArticle(id int) (Article, int) {
	var article Article
	err := db.Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return article, errormessage.ERROR_ART_NOT_EXIST
	}
	return article, errormessage.SUCCESS
}

// GetArticle 查询文章列表 分页功能
func GetArticle(pageSize int, pageNum int) ([]Article, int) {
	var articles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles).Error
	if err != nil {
		return nil, errormessage.ERROR
	}
	return articles, errormessage.SUCCESS
}

// EditArticle 编辑文章
func EditArticle(id int, article *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["description"] = article.Description
	maps["content"] = article.Content
	maps["image"] = article.Image

	err := db.Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormessage.ERROR
	}
	return errormessage.SUCCESS
}
