package dto_book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Full struct {
	ID   primitive.ObjectID
	Name string
	Page int
}
