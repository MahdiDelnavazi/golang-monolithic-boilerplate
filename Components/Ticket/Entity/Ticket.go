package Entity

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Like      bool
	Subject   string
	Message   string
	Image     string
	CreatedAt time.Time
}
