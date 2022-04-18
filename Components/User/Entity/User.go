package Entity

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID
	UserName  string
	Password  string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
