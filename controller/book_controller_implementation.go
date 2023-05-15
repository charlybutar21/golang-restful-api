package controller

import (
	"charlybutar21/golang-restful-api/helper"
	"charlybutar21/golang-restful-api/model/dto"
	"charlybutar21/golang-restful-api/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	createBookRequest := dto.CreateBookRequest{}
	helper.ReadFromRequestBody(request, &createBookRequest)

	bookResponse := controller.BookService.Create(request.Context(), createBookRequest)
	webResponse := dto.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	updateBookRequest := dto.UpdateBookRequest{}
	helper.ReadFromRequestBody(request, &updateBookRequest)

	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	updateBookRequest.Id = id
	bookResponse := controller.BookService.Update(request.Context(), updateBookRequest)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	controller.BookService.Delete(request.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookId := params.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	helper.PanicIfError(err)

	bookResponse := controller.BookService.FindById(request.Context(), id)
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	booksResponse := controller.BookService.FindAll(request.Context())
	webResponse := dto.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   booksResponse,
	}

	helper.WriteToResponseBody(writer, &webResponse)
}
