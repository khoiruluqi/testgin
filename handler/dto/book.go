package dto

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Price int    `json:"price" binding:"required,number"`
	AuthorId int `json:"author_id" binding:"required,number"`
}