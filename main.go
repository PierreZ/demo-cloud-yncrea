package main

import (
	"net/http"

	"github.com/PierreZ/demo-cloud-yncrea/controllers"
	"github.com/PierreZ/demo-cloud-yncrea/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	err := models.ConnectDatabaseWithEnvVars()
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
