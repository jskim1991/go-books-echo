package controller

import (
	"books-app/dto"
	"books-app/model"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindAll() ([]model.Book, error) {
	args := m.Called()
	return args[0].([]model.Book), args.Error(1)
}

func TestGetAllBooks(t *testing.T) {
	t.Run("should return 200 status ok", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := MockRepository{}
		mockRepository.On("FindAll").Return([]model.Book{}, nil)

		controller := Controller{&mockRepository}
		controller.GetAllBooks(c)

		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return books", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := MockRepository{}
		mockRepository.On("FindAll").Return([]model.Book{
			{
				Model:  gorm.Model{ID: 1},
				Isbn:   "999",
				Title:  "Learn Something",
				Author: "Jay",
			},
		}, nil)

		controller := Controller{&mockRepository}
		controller.GetAllBooks(c)

		var books []dto.Book
		json.Unmarshal(rec.Body.Bytes(), &books)
		assert.Equal(t, 1, len(books))
		assert.Equal(t, "999", books[0].Isbn)
		assert.Equal(t, "Learn Something", books[0].Title)
		assert.Equal(t, "Jay", books[0].Author)
	})

	t.Run("should call repository to fetch books", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api/books", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		mockRepository := MockRepository{}
		mockRepository.On("FindAll").Return([]model.Book{}, nil)

		controller := Controller{&mockRepository}
		controller.GetAllBooks(c)

		mockRepository.AssertExpectations(t)
	})
}
