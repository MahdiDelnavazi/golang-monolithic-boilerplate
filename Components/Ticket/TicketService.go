package Ticket

import (
	"go.uber.org/zap"
	"golang_monolithic_bilerplate/Common/Validator"
	Ticket "golang_monolithic_bilerplate/Components/Ticket/Request"
	"golang_monolithic_bilerplate/Components/Ticket/Response"
	Controller "golang_monolithic_bilerplate/Components/User"
	"golang_monolithic_bilerplate/Components/User/Request"
)

type TicketService struct {
	ticketRepository *TicketRepository
	userService      *Controller.UserService
	logger           *zap.SugaredLogger
}

func NewTicketService(logger *zap.SugaredLogger, userService *Controller.UserService, ticketRepository *TicketRepository) *TicketService {
	return &TicketService{logger: logger, userService: userService, ticketRepository: ticketRepository}
}

func (ticketService TicketService) CreateTicket(createTicketRequest Ticket.CreateTicketRequest) (Response.CreateTicketResponse, error) {
	// validate username len and not empty
	validationError := Validator.ValidationCheck(createTicketRequest)

	if validationError != nil {
		return Response.CreateTicketResponse{}, validationError
	}
	user, userError := ticketService.userService.GetUser(Request.GetUserRequest{UserName: createTicketRequest.UserName})
	if userError != nil {
		return Response.CreateTicketResponse{}, userError
	}

	createTicketRequest.UserId = user.UserId
	ticket, ticketError := ticketService.ticketRepository.Create(createTicketRequest)
	if ticketError != nil {
		return Response.CreateTicketResponse{}, ticketError
	}
	// we need a transformer
	return Response.CreateTicketResponse{UserId: ticket.UserId, Subject: ticket.Subject, Message: ticket.Message}, nil
}
