package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manoj8861/go-bookstore/pkg/models"
	"github.com/manoj8861/go-bookstore/pkg/utils"
)

var NewBook models.Book

func CreateBook(wr http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	CreatedBook := createBook.CreateBook()
	res, _ := json.Marshal(CreatedBook)
	wr.Header().Set("Conten-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func GetBooks(wr http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		fmt.Println("error while converting to JSON")
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func GetBookById(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	if Id, err := strconv.ParseInt(BookId, 0, 0); err != nil {
		fmt.Println("Error while Parsing ID")
	} else {
		bookdetails, _ := models.GetBookById(Id)
		res, _ := json.Marshal(bookdetails)
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusOK)
		wr.Write(res)
	}
}

func UpdateBook(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	Id, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	wr.Header().Set("Content-Type", "application/json")

	book, db := models.GetBookById(Id)
	if db.RowsAffected == 0 {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	newBook := &models.Book{}
	utils.ParseBody(r, newBook)

	if newBook.Name != "" {
		book.Name = newBook.Name
	}
	if newBook.Author != "" {
		book.Author = newBook.Author
	}
	if newBook.Publication != "" {
		book.Publication = newBook.Publication
	}
	db.Save(&book)

	res, _ := json.Marshal(book)
	wr.WriteHeader(http.StatusOK)
	wr.Write(res)
}

func DeleteBook(wr http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["BookId"]
	Id, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error while parsing BookId")
		return
	}

	book, db := models.GetBookById(Id)
	if db.RowsAffected == 0 {
		wr.WriteHeader(http.StatusNotFound)
		return
	}
	DeletedBook := models.DeleteBook(int64(book.ID))
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(DeletedBook)
	wr.Write(res)
}
