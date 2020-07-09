package dto_book

type List struct {
	Count int64 `form:"count"`
	Page  int64 `form:"page"`
}
