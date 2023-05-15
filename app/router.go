package app

import (
	"charlybutar21/golang-restful-api/controller"
	"charlybutar21/golang-restful-api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(bookController controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
