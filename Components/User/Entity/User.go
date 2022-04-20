package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserName  string             `bson:"UserName"`
	Password  string             `bson:"Password""`
	Active    bool               `bson:"Active"`
	CreatedAt time.Time          `bson:"CreatedAt"`
	UpdatedAt *time.Time         `bson:"UpdatedAt"`
	DeletedAt *time.Time         `bson:"DeletedAt"`
}

//func SetUser(user User) {
//	if user.UpdatedAt != nil {
//		user.UpdatedAt = time.Time{}
//	}
//}
