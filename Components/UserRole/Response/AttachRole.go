package Response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type AttachRole struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"UserName"`
	IsActive  bool               `bson:"IsActive" `
	RoleID    primitive.ObjectID `bson:"RoleId"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	UpdatedAt *time.Time         `bson:"UpdatedAt"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
}
