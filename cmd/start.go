package cmd

import (
	"jettster/middleware"
	"jettster/provider/config"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/spf13/cobra"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	bookController "jettster/controller/book"
	bozController "jettster/controller/boz"
	pingController "jettster/controller/ping"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
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
		router.GET("/ping", pingController.Ping)

		router.POST("/book", bookController.Create)
		router.GET("/book", bookController.List)
		router.GET("/book/:id", bookController.Single)

		router.GET("/boz/ping", bozController.Ping)

		_ = router.Run(":" + PORT)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

}
