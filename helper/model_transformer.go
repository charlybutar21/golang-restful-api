package helper

import (
	"charlybutar21/golang-restful-api/model/domain"
	"charlybutar21/golang-restful-api/model/dto"
)

func ToBookResponse(book domain.Book) dto.BookResponse {
	return dto.BookResponse{
		Id:   book.Id,
		Name: book.Name,
	}
}

func ToBooksResponse(books []domain.Book) []dto.BookResponse {

	var booksResponse []dto.BookResponse
	for _, book := range books {
		booksResponse = append(booksResponse, ToBookResponse(book))
	}
	return booksResponse
}
