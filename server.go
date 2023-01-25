package main

import (
	"books-app/controller"
	"books-app/repository"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	g := e.Group("/api")

	controller := controller.Controller{Repository: &repository.DefaultRepository{}}
	g.GET("/books", controller.GetAllBooks)

	e.Logger.Fatal(e.Start(":8080"))
}
