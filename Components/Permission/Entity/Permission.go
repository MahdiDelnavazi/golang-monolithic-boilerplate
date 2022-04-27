package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"Name"`
}
