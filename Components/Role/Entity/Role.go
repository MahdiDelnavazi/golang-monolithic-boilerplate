package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name int                `bson:"Name"`
}
