package models

import 	"go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID    primitive.ObjectID `json:"id"`
    Total float64 `json:"total"`
}