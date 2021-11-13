package route

import testController "jettster/controller/test"

func (t *T) TestRoutes() {
	t.router.GET("/test/:id", testController.Ping)
}
