package main

import (
	"GinApiGormMysqlElif/config"
	"GinApiGormMysqlElif/controller"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db         *gorm.DB                  = config.SetupDatabaseConnection()
	authRoutes controller.AuthController = controller.NewAuthController()
)

func main() {

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login")
		authRoutes.POST("/register")
	}
	r.Run()

}
