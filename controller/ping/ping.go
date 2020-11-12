package pingController

import (
	"jettster/provider/rabbitmq"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	rabbitmq.CreatePublisher("ginTestExchange", "direct", true, "ginTestQueue")
}

func Ping(c *gin.Context) {
	rabbitmq.Publish("ginTestExchange", "ginTestQueue", "hi from gin test")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
