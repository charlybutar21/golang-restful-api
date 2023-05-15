package repository

import (
	"charlybutar21/golang-restful-api/helper"
	"charlybutar21/golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
}

func NewBookRepository() BookRepository {
	return &BookRepositoryImpl{}
}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "insert into books(name) values (?)"
	result, err := tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = int(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	SQL := "update books set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	SQL := "select id, name from books where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)
	defer rows.Close()

	book := domain.Book{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	SQL := "select id, name from books"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var books []domain.Book
	for rows.Next() {
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		books = append(books, book)
	}
	return books
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	SQL := "delete from books where id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.PanicIfError(err)
}
