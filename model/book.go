package model_book

import (
	"fmt"

	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func new(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}

func Create(name string, page int) *Book {
	book := new(name, page)
	err := mgm.Coll(book).Create(book)
	if err != nil {
		fmt.Println(err)
	}
	return book
}

func Get(id string) (Book, error) {
	result := Book{}
	error := mgm.Coll(&Book{}).FindByID(id, &result)

	if error != nil {
		return result, error
	}
	return result, nil
}

func List(count int64, skip int64) []Book {
	result := []Book{}
	findOptions := options.Find()
	findOptions.SetSort(bson.M{"_id": -1})
	findOptions.SetLimit(count)
	findOptions.SetSkip(skip)

	_ = mgm.Coll(&Book{}).SimpleFind(&result, bson.M{}, findOptions)
	return result
}
