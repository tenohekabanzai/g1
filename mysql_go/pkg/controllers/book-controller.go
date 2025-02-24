package controllers

import (
	"encoding/json"
	"fmt"
	"io"
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
	newBook := &models.Book{}
	ParseBody(r, newBook)
	b := newBook.CreateBook()
	json.NewEncoder(w).Encode(b)

}
var GetBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	newBook, _ := models.GetBookById(id)
	json.NewEncoder(w).Encode(newBook)

}
var GetBooks = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	newBooks := models.GetAllBooks()
	json.NewEncoder(w).Encode(newBooks)

}

var DeleteBook = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content:Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	b := models.DeleteBook(id)
	json.NewEncoder(w).Encode(b)

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

	bookDetails, db := models.GetBookById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	json.NewEncoder(w).Encode(bookDetails)
}
