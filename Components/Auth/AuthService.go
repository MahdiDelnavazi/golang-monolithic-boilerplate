package Controller

import (
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	User "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	"time"
)

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (authUserService AuthService) CreateAccessToken(accessTokenReq User.AccessTokenRequest) (response string, err error) {

	//todo: move to service
	payload, err := token.MakerPaseto.VerifyToken(accessTokenReq.AccessToken)
	if err != nil {
		return "", err
	}

	// todo: change token duration
	newToken, err := token.MakerPaseto.CreateToken(payload.Username, time.Hour*10000)
	if err != nil {
		return "", err
	}

	return newToken, err
}
