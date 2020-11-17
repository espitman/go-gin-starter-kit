package pingController

import (
	"jettster/provider/elk"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

func Ping(c *gin.Context) {
	elk.CreateIndex("tweets")
	tweet := Tweet{User: "saeed5", Message: "Hi5"}
	result := elk.AddData("tweets", "doc", tweet)
	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
