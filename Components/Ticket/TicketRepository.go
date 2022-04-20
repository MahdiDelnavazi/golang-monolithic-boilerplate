package Ticket

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang_monolithic_bilerplate/Common/Config"
	Entity2 "golang_monolithic_bilerplate/Components/Ticket/Entity"
	Ticket2 "golang_monolithic_bilerplate/Components/Ticket/Request"
	"time"
)

type TicketRepository struct {
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (TicketRepository *TicketRepository) Create(request Ticket2.CreateTicketRequest) (Entity2.TicketMongo, error) {
	ticket := Entity2.TicketMongo{}
	//queryError := Config.DB.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
	//	request.UserId, request.Subject, request.Message, request.Image, request.Like)

	result, queryError := Config.TicketCollection.InsertOne(Config.DBCtx, Entity2.TicketMongo{ID: primitive.NewObjectID(), UserId: request.UserId,
		Subject: request.Subject, Message: request.Message, CreatedAt: time.Now()})

	if queryError != nil {
		return Entity2.TicketMongo{}, nil
	}
	if err := Config.TicketCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&ticket); err != nil {
		return Entity2.TicketMongo{}, err
	}
	return ticket, queryError

}
