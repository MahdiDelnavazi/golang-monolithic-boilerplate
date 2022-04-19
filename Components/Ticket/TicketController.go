package Ticket

import (
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Helper"
	"golang_monolithic_bilerplate/Common/Response"
	Ticket "golang_monolithic_bilerplate/Components/Ticket/Request"
	Response2 "golang_monolithic_bilerplate/Components/Ticket/Response"

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

	//fmt.Println("before call controler", ticketRequest)
	ticketResponse, responseError := ticketControler.ticketService.CreateTicket(ticketRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response2.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "ticket have been created", Data: Ticket.CreateTicketRequest{UserName: ticketResponse.UserName, Message: ticketResponse.Message, Subject: ticketResponse.Subject}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
