package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Ingredient struct {
	Id        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"Name"`
	CreatedAt time.Time          `bson:"CreatedAt"`
}
