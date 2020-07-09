package dto_book

import "go.mongodb.org/mongo-driver/bson/primitive"

type Full struct {
	ID         primitive.ObjectID `json:"_id"`
	Name       string
	Page       int
	Created_at string
	Updated_at string
}
