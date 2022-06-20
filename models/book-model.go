package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ziauddinmohammad/basic-go-CRUD-sql-service/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

type Book_Response struct {
	Data    *Book  `json:"data,omitempty"`
	Message string `json:"message"`
}

func DbInit() {
	db = config.Connect()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Raw("select * from books where deleted_at is null").Scan(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var GetBook Book
	db := db.Where("id=?", Id).Find(&GetBook)
	return &GetBook, db
}

func BookExists(Id int64) bool {
	bookcheck, _ := GetBookById(Id)
	return bookcheck.ID != 0
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Find(&book)
	db.Delete(&book)
	return book
}
