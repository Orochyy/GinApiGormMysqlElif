package main

import (
	"GinApiGormMysqlElif/config"
	"GinApiGormMysqlElif/controller"
	"GinApiGormMysqlElif/middleware"
	"GinApiGormMysqlElif/migrations"
	"GinApiGormMysqlElif/repository"
	"GinApiGormMysqlElif/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

var (
	db                *gorm.DB                     = config.SetupDatabaseConnection()
	userRepository    repository.UserRepository    = repository.NewUserRepository(db)
	bookRepository    repository.BookRepository    = repository.NewBookRepository(db)
	bankRepository    repository.BankRepository    = repository.NewBankRepository(db)
	articleRepository repository.ArticleRepository = repository.NewArticleRepository(db)
	jwtService        service.JWTService           = service.NewJWTService()
	userService       service.UserService          = service.NewUserService(userRepository)
	bookService       service.BookService          = service.NewBookService(bookRepository)
	bankService       service.BankService          = service.NewBankService(bankRepository)
	articleService    service.ArticleService       = service.NewArticleService(articleRepository)
	authService       service.AuthService          = service.NewAuthService(userRepository)
	authController    controller.AuthController    = controller.NewAuthController(authService, jwtService)
	userController    controller.UserController    = controller.NewUserController(userService, jwtService)
	bookController    controller.BookController    = controller.NewBookController(bookService, jwtService)
	bankController    controller.BankController    = controller.NewBankController(bankService, jwtService)
	articleController controller.ArticleController = controller.NewArticleController(articleService, jwtService)
	Migrations                                     = migrations.DbMigrate
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := r.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := r.Group("api/books", middleware.AuthorizeJWT(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindByID)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	articleRoutes := r.Group("api/articles", middleware.AuthorizeJWT(jwtService))
	{
		articleRoutes.GET("/", articleController.All)
		articleRoutes.POST("/", articleController.Insert)
		articleRoutes.GET("/:id", articleController.FindByID)
		articleRoutes.PUT("/:id", articleController.Update)
		articleRoutes.DELETE("/:id", articleController.Delete)
	}
	bankRoutes := r.Group("api/bank", middleware.AuthorizeJWT(jwtService))
	{
		bankRoutes.GET("/", bankController.All)
		bankRoutes.POST("/", bankController.Insert)
		bankRoutes.GET("/:id", bankController.FindByID)
		bankRoutes.PUT("/:id", bankController.Update)
		bankRoutes.DELETE("/:id", bankController.Delete)
	}
	go Migrations()

	r.Run(":8080")
}
