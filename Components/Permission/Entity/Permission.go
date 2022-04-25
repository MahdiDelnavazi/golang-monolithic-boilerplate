package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Permission struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name int                `bson:"Name"`
}
