package services

import (
	"gin-test/core/domains"
	"gin-test/handler/dto"
	"gin-test/repositories"
)

type BookService interface {
	FindAll() ([]domains.Book, error)
	FindById(ID int) (*domains.Book, error)
	Create(bookRequest dto.BookInput) (domains.Book, error)
	Update(ID int, bookRequest dto.BookInput) (*domains.Book, error)
	Delete(ID int) (*domains.Book, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(repository repositories.BookRepository) *bookService {
	return &bookService{repository}
}

func (s *bookService) FindAll() ([]domains.Book, error) {
	books, err := s.bookRepo.FindAll()

	return books, err
}

func (s *bookService) FindById(ID int) (*domains.Book, error) {
	book, err := s.bookRepo.FindById(ID)

	return book, err
}

func (s *bookService) Create(bookRequest dto.BookInput) (domains.Book, error) {
	book := domains.Book{
		Title: bookRequest.Title,
		Price: bookRequest.Price,
		AuthorID: bookRequest.AuthorId,
	}

	newBook, err := s.bookRepo.Create(book)

	return newBook, err
}

func (s *bookService) Update(ID int, bookRequest dto.BookInput) (*domains.Book, error) {
	book, err := s.bookRepo.FindById(ID)

	if err != nil {
		return book, err
	}

	book.Title = bookRequest.Title
	book.Price = bookRequest.Price
	book.AuthorID = bookRequest.AuthorId

	newBook, err := s.bookRepo.Update(*book)

	return &newBook, err
}

func (s *bookService) Delete(ID int) (*domains.Book, error) {
	book, err := s.bookRepo.FindById(ID)

	if err != nil {
		return book, err
	}

	newBook, err := s.bookRepo.Delete(*book)

	return &newBook, err
}