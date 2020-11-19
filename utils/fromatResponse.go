package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatResponse(c *gin.Context, response gin.H) {

	c.JSON(http.StatusOK, gin.H{
		"payload": response,
		"status":  http.StatusOK,
		"message": "ok",
	})
}
