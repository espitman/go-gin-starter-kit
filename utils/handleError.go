package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error":  err.Error(),
		"status": http.StatusBadRequest,
	})
}
