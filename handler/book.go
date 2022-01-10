package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ivanrizkys/go-rest-example/book"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"nama": "Ivan Rizky Saputra",
		"kota": "Semarang",
	})
}

func BooksHandler(c *gin.Context) {
	slug := c.Param("slug")

	c.JSON(http.StatusOK, gin.H{
		"Message": "Get Book succesfully",
		"slug":    slug,
	})
}

func QueryHandler(c *gin.Context) {
	// * get request query
	tittle := c.Query("tittle")
	id := c.Query("id")
	// * get request header
	header := c.GetHeader("Content-Type")
	fmt.Println(tittle)
	fmt.Println(id)
	fmt.Println(header)
	c.JSON(http.StatusOK, gin.H{
		"Messsage": "Get query succesfully",
	})
}

// * handler post request
func PostHandler(c *gin.Context) {
	var bookInput book.Input
	err := c.ShouldBindJSON(&bookInput)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bookInput.Title)
	fmt.Println(bookInput.Price)
	c.JSON(http.StatusOK, gin.H{
		"Message":   "Post succesfully",
		"title":     bookInput.Title,
		"price":     bookInput.Price,
		"sub_title": bookInput.SubTittle,
	})
}

func handleError(c *gin.Context) {
	message := recover()
	if message != nil {
		c.JSON(http.StatusBadGateway, "Internal Server Error")
	}
}

func WithValidator(c *gin.Context) {
	var inputBooks book.Validator
	err := c.ShouldBindJSON(&inputBooks)
	if err != nil {
		// * jika error maka akan menampilkan response status error
		defer handleError(c)
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post with validator request body succesfully",
	})
}
