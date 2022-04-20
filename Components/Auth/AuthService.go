package Controller

import (
	token "golang_monolithic_bilerplate/Common/Token"
	User2 "golang_monolithic_bilerplate/Components/Auth/Request"
	"time"
)

type AuthService struct {
}

func NewAuthUserService() *AuthService {
	return &AuthService{}
}

func (authUserService AuthService) CreateAccessToken(accessTokenReq User2.AccessTokenRequest) (response string, err error) {

	//todo: move to service
	payload, err := token.MakerPaseto.VerifyToken(accessTokenReq.AccessToken)
	if err != nil {
		return "", err
	}
	newToken, err := token.MakerPaseto.CreateToken(payload.Username, time.Hour)
	if err != nil {
		return "", err
	}

	return newToken, err
}
