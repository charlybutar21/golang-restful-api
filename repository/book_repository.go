package repository

import (
	"charlybutar21/golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	Update(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Book
	Delete(ctx context.Context, tx *sql.Tx, book domain.Book)
}
