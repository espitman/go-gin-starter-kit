package bookController

import (
	dto_book_request "jettster/dto/book/request"
	model_book "jettster/model/book"
	"jettster/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Create Book
// @Summary Create Book
// @ID create-book
// @Accept  json
// @Param body body dto_book_request.Create true "the request body"
// @Success 200 {object} dto_book_response.Full
// @Failure 400 {object} dto_error.Error
// @Router /book [post]
func Create(c *gin.Context) {
	var body dto_book_request.Create
	err := c.BindJSON(&body)
	if err != nil {
		utils.HandleErrorBadRequest(c, err)
		return
	}
	book := model_book.Create(body.Name, body.Page)
	utils.FormatResponse(c, gin.H{
		"book": book,
	})
}

// List of Book
// @Summary List of Books
// @Param query query dto_book_request.List true "query params"
// @Success 200 {object} dto_book_response.Summary
// @Failure 400 {object} dto_error.Error
// @Router /book [get]
func List(c *gin.Context) {
	var query dto_book_request.List
	err := c.BindQuery(&query)
	if err != nil {
		utils.HandleErrorBadRequest(c, err)
		return
	}
	var count int64 = 10
	var page int64 = 1
	if query.Count != "" {
		count, _ = strconv.ParseInt(query.Count, 10, 64)
	}
	if query.Page != "" {
		page, _ = strconv.ParseInt(query.Page, 10, 64)
	}
	skip := (page - 1) * count
	books := model_book.List(count, skip)
	utils.FormatResponse(c, gin.H{
		"books": books,
	})
}

// Single Book
// @Summary Single Book
// @Accept  json
// @Param id path string true "url params"
// @Success 200 {object} dto_book_response.Full
// @Failure 400 {object} dto_error.Error
// @Failure 404 {object} dto_error.Error
// @Router /book/{id} [get]
func Single(c *gin.Context) {
	var params dto_book_request.Details
	err := c.BindUri(&params)
	if err != nil {
		utils.HandleErrorBadRequest(c, err)
		return
	}
	book, err := model_book.Get(params.ID)
	if err == nil {
		utils.FormatResponse(c, gin.H{
			"book": book,
		})
		return
	}
	utils.HandleErrorNotFound(c, err)
}
