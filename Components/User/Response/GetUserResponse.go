package Response

import "github.com/google/uuid"

type GetUserResponse struct {
	UserId   uuid.UUID `json:"subject" validate:"required"`
	UserName string    `json:"username" validate:"required,min=3"`
}
