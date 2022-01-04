package repository

import (
	"GinApiGormMysqlElif/entity"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	InsertArticle(a entity.Articles) entity.Articles
	UpdateArticle(a entity.Articles) entity.Articles
	DeleteArticle(a entity.Articles) *gorm.DB
	AllArticle() []entity.Articles
	FindArticleByID(articlesID uint64) entity.Articles
}

type articleConnection struct {
	connection *gorm.DB
}

func NewArticleRepository(dbConn *gorm.DB) ArticleRepository {
	return &articleConnection{
		connection: dbConn,
	}
}

func (db *articleConnection) InsertArticle(a entity.Articles) entity.Articles {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *articleConnection) UpdateArticle(a entity.Articles) entity.Articles {
	db.connection.Save(&a)
	db.connection.Preload("User").Find(&a)
	return a
}

func (db *articleConnection) DeleteArticle(a entity.Articles) *gorm.DB {
	//db.connection.Delete(&a)
	return db.connection.Delete(&a)
}

func (db *articleConnection) FindArticleByID(articlesID uint64) entity.Articles {
	var articles entity.Articles
	db.connection.Preload("User").Find(&articles, articlesID)
	return articles
}

func (db *articleConnection) AllArticle() []entity.Articles {
	var articles []entity.Articles
	db.connection.Preload("User").Find(&articles)
	return articles
}
