package main

import (
	"log"
	"net/http"

	"github.com/codelikesuraj/gdsc-challenge-day-nine-ten/controllers"
	"github.com/codelikesuraj/gdsc-challenge-day-nine-ten/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	PORT = "3000"
)

func main() {
	// initialize database
	var err error
	if DB, err = gorm.Open(sqlite.Open("bookstoreapi.db"), &gorm.Config{}); err != nil {
		log.Fatalln("Error connecting to database:", err.Error())
	}

	// run migrations
	if err := DB.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatalln("Error running migrations:", err.Error())
	}

	r := gin.Default()

	BookController := controllers.BookController{DB: DB}
	UserController := controllers.UserController{DB: DB}

	Middleware := Middleware{DB: DB}

	r.POST("/register", UserController.Register)
	r.POST("/login", UserController.Login)
	r.POST("/refresh-token", UserController.RefreshToken)

	r.GET("/earnings", Middleware.IsAdmin, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "you are an admin"})
	})

	r.Group("", Middleware.Authenticate).
		GET("/books", BookController.GetAllBooks).
		GET("/books/:id", BookController.GetABook).
		POST("/books", BookController.CreateBook)

	log.Fatalln(r.Run(":" + PORT))
}
