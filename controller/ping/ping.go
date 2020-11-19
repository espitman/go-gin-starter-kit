package pingController

import (
	"jettster/utils"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	utils.FormatResponse(c, gin.H{
		"message": "pong",
	})
}
