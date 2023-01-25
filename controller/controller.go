package controller

import (
	"books-app/dto"
	"books-app/repository"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	repository.Repository
}

func (m *Controller) GetAllBooks(c echo.Context) error {
	fetchedBooks, _ := m.Repository.FindAll()
	var books []dto.Book
	for _, fetchedBook := range fetchedBooks {
		book := dto.Book{
			Isbn:   fetchedBook.Isbn,
			Title:  fetchedBook.Title,
			Author: fetchedBook.Author,
		}
		books = append(books, book)
	}

	return c.JSON(http.StatusOK, books)
}
