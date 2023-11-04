package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/zrqx/Antah/initializers"
	"github.com/zrqx/Antah/models"
)

func BooksCreate(c *gin.Context) {
	// Read Request Data
	var body struct {
		Title       string
		Author      string
		Description string
	}
	c.Bind(&body)

	// Create a Book
	book := models.Book{Title: body.Title, Author: body.Author, Description: body.Description}
	result := initializers.DB.Create(&book)
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return the Data
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksIndex(c *gin.Context) {
	// Query the Books
	var books []models.Book
	initializers.DB.Find(&books)

	// Return the Data
	c.JSON(200, gin.H{
		"books": books,
	})
}

func BooksShow(c *gin.Context) {
	// Read the Request Param
	id := c.Param("id")

	// Query the Book
	var book models.Book
	initializers.DB.First(&book, id)

	// Return the Data
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksUpdate(c *gin.Context) {
	// Read the Request Param
	id := c.Param("id")

	// Get the data off Request Body
	var body struct {
		Title       string
		Author      string
		Description string
	}
	c.Bind(&body)

	// Find the Book
	var book models.Book
	initializers.DB.First(&book, id)

	// Update the Book
	initializers.DB.Model(&book).Updates(models.Book{Title: body.Title, Author: body.Author, Description: body.Description})

	// Return the Data
	c.JSON(200, gin.H{
		"book": book,
	})
}

func BooksDelete(c *gin.Context) {
	// Get the ID of the Book
	id := c.Param("id")

	// Delete the Book
	initializers.DB.Delete(&models.Book{}, id)

	// Return Something
	c.Status(200)
}
