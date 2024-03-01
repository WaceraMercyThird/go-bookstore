package models

import (
	"github.com/jinzhu/gorm"
	"github.com/WaceraMercyThird/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:json: "name"`
	Author      string `gorm:json: "author"`
	Publication string `gorm:json: "pablication"`
}

func init() {	
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b

}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBooks Book
	db := db.Where("ID=?").Find(&getBooks)

	return &getBooks, db
}

func DeleteBook(ID uint64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)

	return book
}
