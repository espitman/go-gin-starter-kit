package pingController

import (
	"jettster/utils"

	"github.com/gin-gonic/gin"
)

// Ping
// @Summary Ping
// @Accept  json
// @Success 200 {object} dto_ping.Full
// @Router /ping [get]
func Ping(c *gin.Context) {
	utils.FormatResponse(c, gin.H{
		"message": "pong",
	})
}
