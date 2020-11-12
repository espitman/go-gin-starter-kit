package route

import (
	pingController "jettster/controller/ping"
)

func (t *T) PingRoutes() {
	t.router.GET("/ping", pingController.Ping)
}
