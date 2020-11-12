package route

import bozController "jettster/controller/boz"

func (t *T) BozRoutes() {
	t.router.GET("/boz/ping", bozController.Ping)
}
