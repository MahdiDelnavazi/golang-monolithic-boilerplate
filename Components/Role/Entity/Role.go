package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Role struct {
	Id            primitive.ObjectID   `bson:"_id"`
	PermissionsId []primitive.ObjectID `bson:"Permissions"`
	Name          string               `bson:"Name"`
	CreatedAt     time.Time            `bson:"CreatedAt"`
	UpdatedAt     *time.Time           `bson:"UpdatedAt"`
	DeletedAt     *time.Time           `bson:"DeletedAt"`
}
