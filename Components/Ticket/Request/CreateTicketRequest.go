package Ticket

type CreateTicketRequest struct {
	Subject string `json:"subject" validate:"required"`
	Message string `json:"message" validate:"required"`
	Like    bool   `json:"like" validate:"required"`
	Image   string `json:"image"`
}
