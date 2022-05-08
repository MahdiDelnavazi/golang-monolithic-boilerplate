package Controller

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	User "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

var service *AuthService

func init() {
	service = NewAuthService()

	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig("../../.test.env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	token.NewPasetoMaker(config.Token.TokenSymmetricKey)
}

func TestAuthService_CreateAccessToken(t *testing.T) {
	require.NotNil(t, service)

	rand.Seed(time.Now().UnixNano())
	username := Helper.RandomString(5)
	token, err := token.MakerPaseto.CreateToken(username, time.Hour)
	require.NoError(t, err)

	CreateAccessReq := User.AccessTokenRequest{AccessToken: token}

	newToken, err := service.CreateAccessToken(CreateAccessReq)
	require.NoError(t, err)
	require.NotEmpty(t, newToken)

}
