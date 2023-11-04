package main

import (
	"github.com/zrqx/Antah/initializers"
	"github.com/zrqx/Antah/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
