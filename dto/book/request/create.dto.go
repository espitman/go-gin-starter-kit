package dto_book_request

type Create struct {
	Name string `json:"name" binding:"required"`
	Page int    `json:"page" binding:"required"`
}
