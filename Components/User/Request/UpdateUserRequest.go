package Request

import (
	"time"
)

type UpdateUserRequest struct {
	ID       string `bson:"_id" json:"id" form:"id" `
	UserName string `bson:"UserName" json:"userName"`
	//Password  string     `bson:"Password" json:"password"`
	Active    bool       `bson:"Active" json:"active"`
	CreatedAt time.Time  `bson:"CreatedAt" json:"createdAt"`
	UpdatedAt *time.Time `bson:"UpdatedAt" json:"updatedAt"`
}

//func SetUser(user User) {
//	if user.UpdatedAt != nil {
//		user.UpdatedAt = time.Time{}
//	}
//}
