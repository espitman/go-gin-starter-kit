package testController

import (
	gin "github.com/gin-gonic/gin"
	utils "jettster/utils"
)

// Single test
// @Summary Single test
// @ID single-test
// @Accept  json
// @Param id path string true "url params"
// @Success 200 {object} dto_test_response.Full
// @Failure 400 {object} dto_error.Error
// @Router /test/{id} [get]
func Ping(c *gin.Context) {
	utils.FormatResponse(c, gin.H{"ID": "test"})
}
