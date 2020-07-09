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
	var req dto_book.Create
	err := c.BindJSON(&req)
	if err != nil {
		utils.HandleError(c, err)
		return
	}
	book := model_book.Create(req.Name, req.Page)
	c.JSON(http.StatusOK, book)
}

// List of Book
// @Summary List of Books
// @Accept  json

// @Success 200 {object} dto_book.Summary
// @Router /book [get]
func ListOfBooks(c *gin.Context) {
	// @Param req body dto_book.List true "the request body"
	// var req dto_book.List
	// err := c.BindJSON(&req)
	// if err != nil {
	// 	utils.HandleError(c, err)
	// 	return
	// }
	books := model_book.List(10, 1)
	c.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}
