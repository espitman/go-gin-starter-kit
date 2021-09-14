package dto_book_request

type List struct {
	Count string `form:"count"`
	Page  string `form:"page"`
}
