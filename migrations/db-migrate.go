package migrations

import (
	"GinApiGormMysqlElif/config"
	"GinApiGormMysqlElif/entity"
	"gorm.io/gorm"
)

var db *gorm.DB = config.SetupDatabaseConnection()

func DbMigrate() {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Book{})
	db.AutoMigrate(&entity.Articles{})
	db.AutoMigrate(&entity.Bank{})
}
