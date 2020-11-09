package main

import (
	"jettster/cmd"
	"jettster/provider/db"

	_ "jettster/docs"
)

func init() {
	db.Connect()
}

// @title Jettster Swagger API
// @version 1.0
// @description This is a Jettster service APIs.
// @BasePath /

func main() {

	cmd.Execute()

}
