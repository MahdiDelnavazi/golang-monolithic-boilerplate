package Controller

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var userService *UserService

func init() {
	Test.OpenTestingDatabase()
	userService = NewUserService(NewUserRepository())

	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig("../../.test.env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	token.NewPasetoMaker(config.Token.TokenSymmetricKey)

}

func TestUserService_ChangeActiveStatus(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)
	require.True(t, userResponse.IsActive)

	changeStatusRequest := Request.ChangeStatusRequest{ID: userResponse.ID.Hex()}
	changeStatusResponse, err := userService.ChangeActiveStatus(changeStatusRequest)
	require.NoError(t, err)
	require.NotEmpty(t, changeStatusResponse)
	require.NotNil(t, changeStatusResponse)
	require.False(t, changeStatusResponse.IsActive)
}

func TestUserService_ChangePassword(t *testing.T) {
	password := Helper.RandomString(8)
	username := Helper.RandomString(5)
	userRequest := Request.CreateUserRequest{UserName: username, Password: password}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	changePassRequest := Request.ChangePasswordRequest{Id: userResponse.ID.Hex(), CurrentPassword: password, NewPassword: Helper.RandomString(8)}
	changePassResponse, err := userService.ChangePassword(changePassRequest)

	require.NoError(t, err)
	require.Equal(t, changePassResponse.ID, userResponse.ID)
}

func TestUserService_Create(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)
}

func TestUserService_GetAllUsers(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	getAllUserResponse, err := userService.GetAllUsers(1, 1)
	require.NoError(t, err)
	require.NotEmpty(t, getAllUserResponse)
	require.NotNil(t, getAllUserResponse)

}

func TestUserService_GetUser(t *testing.T) {
	username := Helper.RandomString(5)
	userRequest := Request.CreateUserRequest{UserName: username, Password: Helper.RandomString(8)}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	getuserRequest := Request.GetUserRequest{UserName: username}
	getAllUserResponse, err := userService.GetUser(getuserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, getAllUserResponse)
	require.NotNil(t, getAllUserResponse)
}

func TestUserService_GetUserById(t *testing.T) {
	username := Helper.RandomString(5)
	userRequest := Request.CreateUserRequest{UserName: username, Password: Helper.RandomString(8)}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	getuserRequest := Request.GetUser{ID: userResponse.ID.Hex()}
	getUserResponse, err := userService.GetUserById(getuserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, getUserResponse)
	require.NotNil(t, getUserResponse)
	require.Equal(t, getUserResponse.UserName, username)
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
	loginUserRespose, err := userService.LoginUser(loginUserRequest)
	require.NoError(t, err)
	require.NotNil(t, loginUserRespose)
	require.Equal(t, loginUserRespose.UserName, username)
	require.NotEmpty(t, loginUserRespose.AccessToken)
	require.NotEmpty(t, loginUserRespose.RefreshToken)

}

func TestUserService_UpdateUser(t *testing.T) {
	username := Helper.RandomString(5)
	password := Helper.RandomString(8)
	userRequest := Request.CreateUserRequest{UserName: username, Password: password}
	userResponse, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, userResponse)
	require.NotNil(t, userResponse)
	require.Equal(t, userResponse.UserName, userRequest.UserName)

	updateUserRequest := Request.UpdateUserRequest{ID: userResponse.ID.Hex(), UserName: Helper.RandomString(5)}
	updateUserRespose, err := userService.UpdateUser(updateUserRequest)
	require.NoError(t, err)
	require.NotNil(t, updateUserRespose)
	require.NotEqual(t, updateUserRespose.UserName, username)

}
