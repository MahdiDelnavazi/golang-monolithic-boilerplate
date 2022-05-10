package Controller

import (
	"errors"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	User "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Components/User/Response"
	"time"
)

type AuthService struct {
	authRepository *AuthRepository
}

func NewAuthService(authRepository *AuthRepository) *AuthService {
	return &AuthService{authRepository: authRepository}
}

func (authService *AuthService) CreateAccessToken(accessTokenReq User.AccessTokenRequest) (response string, err error) {

	payload, err := token.MakerPaseto.VerifyToken(accessTokenReq.AccessToken)
	if err != nil {
		return "", errors.New("access token is not valid")
	}

	// todo: change token duration
	newToken, err := token.MakerPaseto.CreateToken(payload.Username, time.Hour*10000)
	if err != nil {
		return "", err
	}

	return newToken, err
}

func (authService *AuthService) LoginUser(loginUserRequest Request.LoginUserRequest) (Response.LoginUserResponse, error) {
	user, getUserError := authService.authRepository.LoginUser(loginUserRequest)
	if getUserError != nil {
		return Response.LoginUserResponse{}, getUserError
	}

	//create new token for login
	accessToken, err := token.MakerPaseto.CreateToken(loginUserRequest.UserName, time.Hour*10000)
	if err != nil {
		return Response.LoginUserResponse{}, err
	}

	refreshToken, errRefreshToken := token.MakerPaseto.CreateToken(loginUserRequest.UserName, time.Hour*120)
	if errRefreshToken != nil {
		return Response.LoginUserResponse{}, err
	}

	// we need a transformer
	return Response.LoginUserResponse{UserName: user.UserName, AccessToken: accessToken, RefreshToken: refreshToken, Id: user.ID.Hex()}, nil
}
