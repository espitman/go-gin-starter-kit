package cmd

import (
	"jettster/middleware"
	"jettster/provider/config"
	"jettster/route"
	"jettster/worker/consumer"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/spf13/cobra"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Gin",
	Long:  `Start Gin`,
	Run: func(cmd *cobra.Command, args []string) {
		PORT := config.GetString("port")

		consumer.Start()

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
		route.GetRoutes(router)

		_ = router.Run(":" + PORT)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
