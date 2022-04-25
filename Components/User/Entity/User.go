package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"UserName"`
	Password  string             `bson:"Password"`
	IsActive  bool               `bson:"IsActive" `
	Role      string             `bson:"Role"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	UpdatedAt *time.Time         `bson:"UpdatedAt"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
}
