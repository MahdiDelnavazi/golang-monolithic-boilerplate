package Controller

import (
	"golang_monolithic_bilerplate/Common/Config"
	token "golang_monolithic_bilerplate/Common/Token"
	User "golang_monolithic_bilerplate/Components/AuthUser/Request"
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
