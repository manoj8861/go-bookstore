package models

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/manoj8861/go-bookstore/pkg/config"
)

type Book struct {
	gorm.Model
	Name        string `json:"Name"`
	Author      string `json:"Author"`
	Publication string `json:"Publication"`
}

var db *gorm.DB

func init() {
	if err := config.Connect(); err != nil {
		log.Fatal(err)
		return
	}
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	if db.NewRecord(b) {
		db.Create(&b)
	}
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var Book Book
	db := db.Where("Id=?", Id).Find(&Book)
	return &Book, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("Id=?", Id).Delete(book)
	return book
}
