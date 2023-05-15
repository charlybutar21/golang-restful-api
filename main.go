package main

import (
	"charlybutar21/golang-restful-api/app"
	"charlybutar21/golang-restful-api/controller"
	"charlybutar21/golang-restful-api/helper"
	"charlybutar21/golang-restful-api/middleware"
	"charlybutar21/golang-restful-api/repository"
	"charlybutar21/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	bookRepository := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)
	router := app.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
