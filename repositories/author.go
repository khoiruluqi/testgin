package repositories

import (
	"gin-test/core/domains"

	"gorm.io/gorm"
)

type AuthorRepository interface {
	FindAll() ([]domains.Author, error)
	// FindById(ID int) (domains.Author, error)
	// Create(author domains.Author) (domains.Author, error)
}

type authorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *authorRepository {
	return &authorRepository{db}
}

func (r *authorRepository) FindAll() ([]domains.Author, error) {
	var authors []domains.Author

	err := r.db.Preload("Books").Find(&authors).Error

	return authors, err
}
