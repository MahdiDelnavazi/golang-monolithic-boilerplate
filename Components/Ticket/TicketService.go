package Ticket

import (
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	Ticket "github.com/mahdidl/golang_boilerplate/Components/Ticket/Request"
	"github.com/mahdidl/golang_boilerplate/Components/Ticket/Response"
	Controller "github.com/mahdidl/golang_boilerplate/Components/User"
)

type TicketService struct {
	ticketRepository *TicketRepository
	userService      *Controller.UserService
}

func NewTicketService(userService *Controller.UserService, ticketRepository *TicketRepository) *TicketService {
	return &TicketService{userService: userService, ticketRepository: ticketRepository}
}

func (ticketService TicketService) CreateTicket(createTicketRequest Ticket.CreateTicketRequest, userId string) (Response.CreateTicketResponse, error) {
	// validate username len and not empty
	validationError := Validator.ValidationCheck(createTicketRequest)

	if validationError != nil {
		return Response.CreateTicketResponse{}, validationError
	}

	ticket, ticketError := ticketService.ticketRepository.Create(createTicketRequest, userId)
	if ticketError != nil {
		return Response.CreateTicketResponse{}, ticketError
	}
	// we need a transformer
	return Response.CreateTicketResponse{Subject: ticket.Subject, Message: ticket.Message, Image: ticket.Image}, nil
}
