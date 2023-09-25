package main

import (
	"net/http"
	"qmun14/golang-restful-api/app"
	"qmun14/golang-restful-api/controller"
	"qmun14/golang-restful-api/helper"
	"qmun14/golang-restful-api/middleware"
	"qmun14/golang-restful-api/repository"
	"qmun14/golang-restful-api/service"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
