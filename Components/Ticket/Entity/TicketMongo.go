package Entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TicketMongo struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    primitive.ObjectID `bson:"UserId"`
	Like      bool               `bson:"Like"`
	Subject   string             `bson:"Subject"`
	Message   string             `bson:"Message"`
	Image     string             `bson:"Image"`
	CreatedAt time.Time          `bson:"CreatedAt"`
}
