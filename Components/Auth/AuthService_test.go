package Controller

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	User "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	User2 "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

var authService *AuthService
var userService *User2.UserService

func init() {
	authService = NewAuthService(NewAuthRepository())
	userService = User2.NewUserService(User2.NewUserRepository())

	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig("../../.test.env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	token.NewPasetoMaker(config.Token.TokenSymmetricKey)
}

func TestAuthService_CreateAccessToken(t *testing.T) {
	require.NotNil(t, authService)

	rand.Seed(time.Now().UnixNano())
	username := Helper.RandomString(5)
	token, err := token.MakerPaseto.CreateToken(username, time.Hour)
	require.NoError(t, err)

	CreateAccessReq := User.AccessTokenRequest{AccessToken: token}

	newToken, err := authService.CreateAccessToken(CreateAccessReq)
	require.NoError(t, err)
	require.NotEmpty(t, newToken)

}

func TestUserService_LoginUser(t *testing.T) {
	username := Helper.RandomString(5)
	password := Helper.RandomString(8)
	userRequest := Request.CreateUserRequest{UserName: username, Password: password}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	loginUserRequest := Request.LoginUserRequest{UserName: username, Password: password}
	loginUserRespose, err := authService.LoginUser(loginUserRequest)
	require.NoError(t, err)
	require.NotNil(t, loginUserRespose)
	require.Equal(t, loginUserRespose.UserName, username)
	require.NotEmpty(t, loginUserRespose.AccessToken)
	require.NotEmpty(t, loginUserRespose.RefreshToken)

}
