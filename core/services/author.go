package services

import (
	"gin-test/core/domains"
	"gin-test/repositories"
)

type AuthorService interface {
	FindAll() ([]domains.Author, error)
	// FindById(ID int) (domains.Book, error)
	// Create(bookRequest dto.BookInput) (domains.Book, error)
}

type authorService struct {
	authorRepo repositories.AuthorRepository
}

func NewAuthorService(repository repositories.AuthorRepository) *authorService {
	return &authorService{repository}
}

func (s *authorService) FindAll() ([]domains.Author, error) {
	authors, err := s.authorRepo.FindAll()

	return authors, err
}

// func (s *authorService) FindById(ID int) (domains.Book, error) {
// 	book, err := s.bookRepo.FindById(ID)

// 	return book, err
// }

// func (s *authorService) Create(bookRequest dto.BookInput) (domains.Book, error) {
// 	book := domains.Book{
// 		Title: bookRequest.Title,
// 		Price: bookRequest.Price,
// 	}

// 	newBook, err := s.bookRepo.Create(book)

// 	return newBook, err
// }
