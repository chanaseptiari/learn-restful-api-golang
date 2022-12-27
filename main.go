package main

import (
	_ "github.com/go-sql-driver/mysql"

	"chanaseptiari/learn-restful-api-golang/conf"
	"chanaseptiari/learn-restful-api-golang/controller"
	"chanaseptiari/learn-restful-api-golang/helper"
	"chanaseptiari/learn-restful-api-golang/repository"
	"chanaseptiari/learn-restful-api-golang/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := conf.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryConrtoller(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
