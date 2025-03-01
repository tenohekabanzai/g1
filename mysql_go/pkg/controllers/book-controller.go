package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"mysql_go/pkg/config"
	"mysql_go/pkg/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ParseBody(r *http.Request, X interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, X); err != nil {
		return err
	}
	return nil
}

var NewBook models.Book

var CreateBook = func(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content:Type", "application/json")
	db := config.GetDB()
	var newBook models.Book
	ParseBody(r, &newBook)
	db.Create(&newBook)
	json.NewEncoder(w).Encode("Book Added Successfully")

}

var GetBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	params := mux.Vars(r)
	db := config.GetDB()
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	var book models.Book
	db.Find(&book, id)
	json.NewEncoder(w).Encode(book)
}

var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	db := config.GetDB()
	var books []models.Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)

}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	db := config.GetDB()
	db.Delete(&models.Book{}, id)
	json.NewEncoder(w).Encode("Book Deleted Successfully")
}

var UpdateBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	var updateBook = &models.Book{}
	ParseBody(r, updateBook)
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	db := config.GetDB()
	var book models.Book

	db.Find(&book, id)

	if updateBook.Name != "" {
		book.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		book.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		book.Publication = updateBook.Publication
	}

	db.Model(&book).Update(book)
	json.NewEncoder(w).Encode(book)
}
