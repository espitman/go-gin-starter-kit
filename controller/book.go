package controller

import (
	dto_book "jettster/dto/book/request"
	model_book "jettster/model"
	"jettster/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Book
// @Summary Create Book
// @ID create-book
// @Accept  json
// @Param req body dto_book.Create true "the request body"
// @Success 200 {object} dto_book.Full
// @Router /book [post]
func CreateBook(c *gin.Context) {
	var body dto_book.Create
	err := c.BindJSON(&body)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	book := model_book.Create(body.Name, body.Page)
	c.JSON(http.StatusOK, book)
}

// List of Book
// @Summary List of Books
// @Accept  json
// @Success 200 {object} dto_book.Summary
// @Router /book [get]
func ListOfBooks(c *gin.Context) {
	var query dto_book.List
	err := c.Bind(&query)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	skip := (query.Page - 1) * query.Count
	books := model_book.List(query.Count, skip)
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
