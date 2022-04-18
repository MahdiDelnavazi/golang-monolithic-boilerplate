package Response

import "github.com/google/uuid"

type CreateTicketResponse struct {
	UserId   uuid.UUID `json:"subject" validate:"required"`
	UserName string    `json:"username" validate:"required,min=3"`
	Subject  string    `json:"subject" validate:"required"`
	Message  string    `json:"message" validate:"required"`
	Like     bool      `json:"like" validate:"required"`
	Image    []byte    `json:"image"`
}
