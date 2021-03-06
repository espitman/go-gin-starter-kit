package db

import (
	"fmt"
	"jettster/provider/config"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() {
	err := mgm.SetDefaultConfig(nil, config.GetString("db.name"), options.Client().ApplyURI(config.GetString("db.uri")))
	if err != nil {
		fmt.Println(err)
	}
}
