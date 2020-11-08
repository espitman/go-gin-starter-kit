package pingController

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping ...
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
