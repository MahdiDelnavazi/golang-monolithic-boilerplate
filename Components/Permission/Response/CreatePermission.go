package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreatePermission struct {
	ID   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"`
}
