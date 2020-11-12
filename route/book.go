package route

import (
	bookController "jettster/controller/book"
)

func (t *T) BookRoutes() {
	t.router.POST("/book", bookController.Create)
	t.router.GET("/book", bookController.List)
	t.router.GET("/book/:id", bookController.Single)
}
