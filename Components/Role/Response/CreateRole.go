package Response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateRole struct {
	ID   primitive.ObjectID `json:"Id" bson:"_id"`
	Name string             `json:"Name" bson:"Name"  validate:"required"`
}
