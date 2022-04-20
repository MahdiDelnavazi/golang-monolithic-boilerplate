package Entity

import (
	"github.com/google/uuid"
	"time"
)

type Ticket struct {
	UserId    uuid.UUID
	Like      bool
	Subject   string
	Message   string
	Image     string
	CreatedAt time.Time
}
