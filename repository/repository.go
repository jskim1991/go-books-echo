package repository

import (
	"books-app/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]model.Book, error)
}

type DefaultRepository struct {
}

func (m *DefaultRepository) FindAll() ([]model.Book, error) {
	return []model.Book{{
		Model:  gorm.Model{ID: 1},
		Isbn:   "9780321278654",
		Title:  "Extreme Programming Explained: Embrace Change",
		Author: "Kent Beck, Cynthia Andres",
	}}, nil
}
