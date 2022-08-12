package repositories

import (
	"fmt"
	"gin-test/core/domains"

	"gorm.io/gorm"
)

type BookRepository interface {
	FindAll() ([]domains.Book, error)
	FindById(ID int) (*domains.Book, error)
	Create(book domains.Book) (domains.Book, error)
	Update(book domains.Book) (domains.Book, error)
	Delete(book domains.Book) (domains.Book, error)
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

func (r *bookRepository) FindById(ID int) (*domains.Book, error) {
	var book domains.Book

	result := r.db.Find(&book, ID)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &book, nil
}

func (r *bookRepository) Create(book domains.Book) (domains.Book, error) {
	err := r.db.Create(&book).Error

	return book, err
}

func (r *bookRepository) Update(book domains.Book) (domains.Book, error) {
	err := r.db.Save(&book).Error

	return book, err
}

func (r *bookRepository) Delete(book domains.Book) (domains.Book, error) {
	err := r.db.Delete(&book).Error

	return book, err
}