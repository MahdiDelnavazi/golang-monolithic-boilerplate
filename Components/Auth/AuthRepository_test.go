package Controller

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	User "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var authRepository *AuthRepository
var userRepository *User.UserRepository

func init() {
	Test.OpenTestingDatabase()
	authRepository = NewAuthRepository()
	userRepository = User.NewUserRepository()
}

func TestAuthRepository_LoginUser(t *testing.T) {
	password := Helper.RandomString(8)
	username := Helper.RandomString(5)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we read it from db and test login
	creatUserRequest := Request.CreateUserRequest{UserName: username, Password: hashedPassword}

	user, err := userRepository.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	loginRequest := Request.LoginUserRequest{UserName: username, Password: password}
	login, loginErr := authRepository.LoginUser(loginRequest)
	require.NoError(t, loginErr)
	require.NotEmpty(t, login)
	require.NotNil(t, login)
	require.Equal(t, login.UserName, user.UserName)
}
