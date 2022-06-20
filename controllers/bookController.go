package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ziauddinmohammad/basic-go-CRUD-sql-service/models"
	"github.com/ziauddinmohammad/basic-go-CRUD-sql-service/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all books called")

	AllBooks := models.GetAllBooks()
	result, _ := json.Marshal(AllBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get book by id called")

	vars := mux.Vars(r)
	Id_String := vars["bookId"]
	Id, _ := strconv.ParseInt(Id_String, 0, 0)
	book, _ := models.GetBookById(Id)
	w.Header().Set("Content-Type", "application/json")
	if book.ID != 0 {
		result, _ := json.Marshal(book)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	response := models.Book_Response{Message: "id not found"}
	result, _ := json.Marshal(response)
	w.Write(result)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create book called")

	newbook := models.Book{}
	utils.ParseBody(r, &newbook)
	b := newbook.CreateBook()
	response := models.Book_Response{Data: b, Message: "new book added"}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id_String := vars["bookId"]
	Id, _ := strconv.ParseInt(Id_String, 0, 0)
	if !models.BookExists(Id) {
		response := models.Book_Response{Message: "book not found"}
		result, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}
	existing_book, db := models.GetBookById(Id)

	newbook := models.Book{}
	utils.ParseBody(r, &newbook)
	if newbook.Author != "" {
		existing_book.Author = newbook.Author
	}
	if newbook.Name != "" {
		existing_book.Name = newbook.Name
	}
	if newbook.Publication != "" {
		existing_book.Publication = newbook.Publication
	}
	db.Save(&existing_book)
	response := models.Book_Response{Data: existing_book, Message: "book updated"}
	result, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete book called")

	vars := mux.Vars(r)
	Id_String := vars["bookId"]
	Id, _ := strconv.ParseInt(Id_String, 0, 0)
	if models.BookExists(Id) {
		models.DeleteBook(Id)
		response := models.Book_Response{Message: "book deleted"}
		result, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	}
	w.WriteHeader(http.StatusNotFound)
	response := models.Book_Response{Message: "book not found"}
	result, _ := json.Marshal(response)
	w.Write(result)
}
