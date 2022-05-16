package Ticket

import (
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Components/Ticket/Entity"
	Ticket "github.com/mahdidl/golang_boilerplate/Components/Ticket/Request"
	UserEntity "github.com/mahdidl/golang_boilerplate/Components/User/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type TicketRepository struct {
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (TicketRepository *TicketRepository) Create(request Ticket.CreateTicketRequest, userId string) (Entity.Ticket, error) {
	ticket := Entity.Ticket{}
	user := UserEntity.User{}
	//queryError := Config.DB.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
	//	request.UserId, request.Subject, request.Message, request.Image, request.Like)

	primitiveUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Entity.Ticket{}, fmt.Errorf("id is not valid")
	}

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"_id": primitiveUserId}).Decode(&user)
	if queryError != nil {
		return Entity.Ticket{}, fmt.Errorf("user not found")
	}

	result, queryError := Config.TicketCollection.InsertOne(Config.DBContext, Entity.Ticket{ID: primitive.NewObjectID(), UserId: primitiveUserId,
		Subject: request.Subject, Message: request.Message, CreatedAt: time.Now()})

	if queryError != nil {
		return Entity.Ticket{}, nil
	}
	if err := Config.TicketCollection.FindOne(Config.DBContext, bson.M{"_id": result.InsertedID}).Decode(&ticket); err != nil {
		return Entity.Ticket{}, err
	}
	return ticket, queryError

}
