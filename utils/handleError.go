package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleErrorBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  err.Error(),
		"status": http.StatusBadRequest,
	})
}

func HandleErrorNotFound(c *gin.Context, err error) {
	c.JSON(http.StatusNotFound, gin.H{
		"error":  err.Error(),
		"status": http.StatusNotFound,
	})
}
