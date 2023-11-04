package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zrqx/Antah/controllers"
	"github.com/zrqx/Antah/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/books", controllers.BooksCreate)
	r.GET("/books", controllers.BooksIndex)
	r.GET("/books/:id", controllers.BooksShow)
	r.PUT("/books/:id", controllers.BooksUpdate)
	r.DELETE("/books/:id", controllers.BooksDelete)
	r.Run()
}
