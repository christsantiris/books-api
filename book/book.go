package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"books-api/database"
)

type Book struct {
	gorm.Model
	Title string `json:"title"`
	Author string `json:"author"`
	Rating int `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	c.JSON(books)
}
func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}
func NewBook(c *fiber.Ctx) {
db := database.DBConn
	var book Book
	book.Title = "The Hobbit"
	book.Author = "Tolkein"
	book.Rating = 10

	db.Create(&book)
	c.JSON(book)
}
func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No book found for Id")
	}
	db.Delete(&book)
	c.Send("Book successfully deleted")
}
