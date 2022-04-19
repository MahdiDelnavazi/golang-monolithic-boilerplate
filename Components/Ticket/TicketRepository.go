package Ticket

import (
	"golang_monolithic_bilerplate/Common/Config"
	Entity2 "golang_monolithic_bilerplate/Components/Ticket/Entity"
	Ticket2 "golang_monolithic_bilerplate/Components/Ticket/Request"
)

type TicketRepository struct {
}

func NewTicketRepository() *TicketRepository {
	return &TicketRepository{}
}

func (TicketRepository *TicketRepository) Create(request Ticket2.CreateTicketRequest) (Entity2.Ticket, error) {
	ticket := Entity2.Ticket{}
	queryError := Config.PostgresDB.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
		request.UserId, request.Subject, request.Message, request.Image, request.Like)

	if queryError != nil {
		return Entity2.Ticket{}, nil
	}
	return ticket, queryError

}
