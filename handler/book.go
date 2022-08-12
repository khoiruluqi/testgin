package handler

import (
	"gin-test/core/services"
	"gin-test/handler/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookHandler struct {
	bookService services.BookService
}

func NewBookHandler(bookService services.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (h *bookHandler) GetAll(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

// func BooksHandler(c *gin.Context) {
// 	id := c.Param("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id": id,
// 	})
// }

// func QueryHandler(c *gin.Context) {
// 	id := c.Query("id")

// 	c.JSON(http.StatusOK, gin.H{
// 		"id": id,
// 	})
// }

func (h *bookHandler) Create(c *gin.Context) {
	var bookInput dto.BookInput

	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.bookService.Create(bookInput)

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"price": bookInput.Price,
	})
}
