package Ticket

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	Entity2 "golang_monolithic_bilerplate/Components/Ticket/Entity"
	Ticket2 "golang_monolithic_bilerplate/Components/Ticket/Request"
)

type TicketRepository struct {
	logger   *zap.SugaredLogger
	database *sqlx.DB
}

func NewTicketRepository(logger *zap.SugaredLogger, database *sqlx.DB) *TicketRepository {
	return &TicketRepository{
		logger:   logger,
		database: database,
	}
}

func (TicketRepository *TicketRepository) Create(request Ticket2.CreateTicketRequest) (Entity2.Ticket, error) {
	ticket := Entity2.Ticket{}
	queryError := TicketRepository.database.Get(&ticket, `SELECT * FROM newTicket($1 , $2 , $3 , $4, $5)`,
		request.UserId, request.Subject, request.Message, request.Image, request.Like)

	if queryError != nil {
		return Entity2.Ticket{}, nil
	}
	return ticket, queryError

}
