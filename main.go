package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

type customBook struct {
	ID    string `json:"id"`
	Price int    `json:"price"`
}

var books = []book{
	{ID: "1", Title: "Pillars of the earth", Author: "Ken Follett", Price: 699},
	{ID: "2", Title: "World without end", Author: "Ken Follett", Price: 699},
	{ID: "3", Title: "Column of Fire", Author: "Ken Follett", Price: 699},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		books,
	)
}

func createBook(c *gin.Context) {
	body := book{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.IndentedJSON(
		http.StatusCreated,
		&body,
	)
}

func editBook(c *gin.Context) {
	c.IndentedJSON(
		http.StatusOK,
		nil,
	)
}

func deleteBook(c *gin.Context) {
	c.IndentedJSON(
		http.StatusNoContent,
		nil,
	)
}

func main() {
	router := gin.Default()

	// GET
	router.GET("/books", getBooks)

	// POST
	router.POST("/books", createBook)

	// PUT
	router.PUT("/books", editBook)

	// DELETE
	router.DELETE("/books", deleteBook)

	router.Run("localhost:8080")
}
