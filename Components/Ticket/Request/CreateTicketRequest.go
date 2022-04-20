package Ticket

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTicketRequest struct {
	UserId   primitive.ObjectID `json:"userId"`
	UserName string             `json:"username" validate:"required,min=3"`
	Subject  string             `json:"subject" validate:"required"`
	Message  string             `json:"message" validate:"required"`
	Like     bool               `json:"like" validate:"required"`
	Image    string             `json:"image"`
}
