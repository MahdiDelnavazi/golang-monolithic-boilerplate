package Controller

import (
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	User "github.com/mahdidl/golang_boilerplate/Components/AuthUser/Request"
)

type AuthUserRepository struct {
}

func NewAuthUserRepository() *AuthUserRepository {
	return &AuthUserRepository{}
}

func (userRepository *AuthUserRepository) LogOut(logoutReq User.LogoutRequest, payload *token.Payload) error {

	err := Config.Redis.Set(payload.Username, logoutReq.Token, 0).Err()

	if err != nil {
		return nil
	}
	return err
}
