package repositories

import (
	"gin-test/core/domains"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]domains.Book, error)
	FindById(ID int) (domains.Book, error)
	Create(book domains.Book) (domains.Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (r *bookRepository) FindAll() ([]domains.Book, error) {
	var books []domains.Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *bookRepository) FindById(ID int) (domains.Book, error) {
	var book domains.Book

	err := r.db.Find(&book, ID).Error

	return book, err
}

func (r *bookRepository) Create(book domains.Book) (domains.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}
