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

	router.GET("/books", bookHandler.GetAll)
	// router.GET("/books/:id", book.BooksHandler)
	// router.GET("/query", book.QueryHandler)
	router.POST("/books", bookHandler.Create)

	router.Run()
}
