package main

import (
	"jettster/controller"
	"jettster/middleware"
	"jettster/provider/config"
	"jettster/provider/db"
	"time"

	_ "jettster/docs"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	db.Connect()
}

// @title Jettster Swagger API
// @version 1.0
// @description This is a Jettster service APIs.
// @BasePath /

func main() {

	PORT := config.GetString("port")

	router := gin.New()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET,POST,PUT,DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	url := ginSwagger.URL("http://localhost:" + PORT + "/swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Use(middleware.SetJsonHeader)
	router.GET("/ping", controller.Ping)

	router.POST("/book", controller.CreateBook)
	router.GET("/book", controller.ListOfBooks)
	router.GET("/book/:id", controller.SingleBook)
	_ = router.Run(":" + PORT)
}
