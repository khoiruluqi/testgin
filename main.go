package main

import (
	"gin-test/core/services"
	"gin-test/handler"
	"gin-test/repositories"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}

func main() {
	dbURL := goDotEnvVariable("POSTGRES_URL")
	DB := repositories.Init(dbURL)

	bookRepository := repositories.NewBookRepository(DB)
	bookService := services.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	booksRouter := router.Group("books")

	booksRouter.GET("/", bookHandler.GetAll)
	booksRouter.GET("/:id", bookHandler.GetById)
	booksRouter.POST("/", bookHandler.Create)
	booksRouter.PUT("/:id", bookHandler.Update)
	booksRouter.DELETE("/:id", bookHandler.Delete)

	router.Run()
}
