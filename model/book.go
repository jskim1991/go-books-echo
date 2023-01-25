package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Isbn   string
	Title  string
	Author string
}
