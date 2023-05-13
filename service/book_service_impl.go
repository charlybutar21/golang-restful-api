package service

import (
	"context"
	"database/sql"
	"github.com/golang-restful-api/helper"
	"github.com/golang-restful-api/model/domain"
	"github.com/golang-restful-api/model/dto"
	"github.com/golang-restful-api/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
	DB             *sql.DB
}

func NewBookService(bookRepository repository.BookRepository, DB *sql.DB) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		DB:             DB,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request dto.CreateBookRequest) dto.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book := domain.Book{
		Name: request.Name,
	}

	book = service.BookRepository.Save(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Update(ctx context.Context, request dto.UpdateBookRequest) dto.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	book.Name = request.Name

	book = service.BookRepository.Update(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	helper.PanicIfError(err)

	service.BookRepository.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) dto.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, bookId)
	helper.PanicIfError(err)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []dto.BookResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	books := service.BookRepository.FindAll(ctx, tx)

	return helper.ToBooksResponse(books)
}
