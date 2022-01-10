package main

import (
	"github.com/Ivanrizkys/go-rest-example/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", handler.RootHandler)

	// * get routing with query
	router.GET("/query", handler.QueryHandler)

	// * get routing with params
	router.GET("/books/:slug", handler.BooksHandler)

	// * post request
	router.POST("/books", handler.PostHandler)

	router.POST("/buku", handler.WithValidator)

	// * modify port, default is 8080
	router.Run()
}
