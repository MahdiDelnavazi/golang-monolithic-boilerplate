package Ticket

import (
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Components/Ticket/Entity"
	Ticket "github.com/mahdidl/golang_boilerplate/Components/Ticket/Request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"time"
)

type TicketRepository struct {
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (TicketRepository *TicketRepository) Create(request Ticket.CreateTicketRequest) (Entity.Ticket, error) {
	ticket := Entity.Ticket{}
	//queryError := Config.DB.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
	//	request.UserId, request.Subject, request.Message, request.Image, request.Like)

	result, queryError := Config.TicketCollection.InsertOne(Config.DBCtx, Entity.Ticket{ID: primitive.NewObjectID(), UserId: request.UserId,
		Subject: request.Subject, Message: request.Message, CreatedAt: time.Now()})

	if queryError != nil {
		return Entity.Ticket{}, nil
	}
	if err := Config.TicketCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&ticket); err != nil {
		return Entity.Ticket{}, err
	}
	return ticket, queryError

}
