package main

import (
	"fmt"
	"net/http"
	"os"

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
	r.GET("/fatal", func(c *gin.Context) {
		fmt.Println("exiting because fatal")
		os.Exit(1)
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	fmt.Println(r.Run())
}
