package main

import (
	_ "github.com/go-sql-driver/mysql"

	"chanaseptiari/learn-restful-api-golang/conf"
	"chanaseptiari/learn-restful-api-golang/controller"
	"chanaseptiari/learn-restful-api-golang/helper"
	"chanaseptiari/learn-restful-api-golang/middleware"
	"chanaseptiari/learn-restful-api-golang/repository"
	"chanaseptiari/learn-restful-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := conf.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryConrtoller(categoryService)

	router := conf.NewRouter(categoryController)

	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
