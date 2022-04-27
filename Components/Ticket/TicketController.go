package Ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	Ticket "github.com/mahdidl/golang_boilerplate/Components/Ticket/Request"
	Response "github.com/mahdidl/golang_boilerplate/Components/Ticket/Response"

	"log"

	"net/http"
)

type TicketController struct {
	ticketService *TicketService
}

func NewTicketController(ticketService *TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (ticketControler *TicketController) CreateTicket(context *gin.Context) {
	var ticketRequest Ticket.CreateTicketRequest
	Helper.Decode(context.Request, &ticketRequest)

	validationError := Validator.ValidationCheck(ticketRequest)
	log.Println(validationError)
	if validationError != nil {
		response := General.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
	}

	ticketResponse, responseError := ticketControler.ticketService.CreateTicket(ticketRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := General.GeneralResponse{Error: false, Message: "ticket have been created", Data: Response.CreateTicketResponse{Message: ticketResponse.Message, Subject: ticketResponse.Subject, Image: ticketResponse.Image}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
