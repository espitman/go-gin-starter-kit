package model_book

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Kamva/mgm"
)

type book struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func new(name string, pages int) *book {
	return &book{
		Name:  name,
		Pages: pages,
	}
}

func Create(name string, page int) primitive.ObjectID {
	book := new(name, page)
	err := mgm.Coll(book).Create(book)
	if err != nil {
		fmt.Println(err)
	}
	return book.ID
}
