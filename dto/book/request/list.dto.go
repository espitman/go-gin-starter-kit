package dto_book

type List struct {
	Count int `json:"count" binding:"required"`
	Page  int `json:"page" binding:"required"`
}
