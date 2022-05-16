package Controller

import (
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/User/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"go.mongodb.org/mongo-driver/bson"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

// LoginUser for login users
func (authRepository *AuthRepository) LoginUser(loginUserRequest Request.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"UserName": loginUserRequest.UserName}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}

	if !Helper.CheckPasswordHash(loginUserRequest.Password, user.Password) {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	return user, queryError
}
