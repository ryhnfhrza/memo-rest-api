package main

import (
	"memoAPI/app"
	"memoAPI/controller"
	"memoAPI/exception"
	"memoAPI/helper"
	"memoAPI/middleware"
	"memoAPI/repository"
	"memoAPI/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	DB := app.GetConnection()
	validate := validator.New()
	memoesRepository := repository.NewMemoesRepository()
	memoesService := service.NewMemoesService(memoesRepository,DB,validate)
	memoesController := controller.NewMemoesController(memoesService)

	router := httprouter.New()

	router.GET("/api/memoes",memoesController.FindAll)
	router.GET("/api/memoes/order-by-title",memoesController.OrderByTitleAsc)
	router.GET("/api/memoes/order-by-id",memoesController.OrderByIdDesc)
	router.GET("/api/memo/:memoesId",memoesController.FindById)
	router.GET("/api/memo-title/:memoesTitle",memoesController.FindByTitle)
	router.PUT("/api/memo/:memoesId",memoesController.Update)
	router.DELETE("/api/memo/:memoesId",memoesController.Delete)
	router.POST("/api/memoes",memoesController.Create)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),	
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}