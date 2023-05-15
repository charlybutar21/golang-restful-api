package service

import (
	"charlybutar21/golang-restful-api/model/dto"
	"context"
)

type BookService interface {
	Create(ctx context.Context, request dto.CreateBookRequest) dto.BookResponse
	Update(ctx context.Context, request dto.UpdateBookRequest) dto.BookResponse
	Delete(ctx context.Context, bookId int)
	FindById(ctx context.Context, bookId int) dto.BookResponse
	FindAll(ctx context.Context) []dto.BookResponse
}
