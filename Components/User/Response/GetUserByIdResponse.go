package Response

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUserByIdResponse struct {
	UserId   primitive.ObjectID `json:"subject" validate:"required"`
	UserName string             `json:"username" validate:"required,min=3"`
}
