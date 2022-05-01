package Controller

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

func init() {
	Test.OpenTestingDatabase()
}

func TestUserRepository_ChangeActiveStatus(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	repo := NewUserRepository()
	require.NotNil(t, repo)

	user, err := repo.CreateUser(creatUserRequest, creatUserRequest.Password)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)
}

func TestUserRepository_ChangePassword(t *testing.T) {

}

func TestUserRepository_CheckUserName(t *testing.T) {

}

func TestUserRepository_CreateUser(t *testing.T) {

}

func TestUserRepository_GetAllUsers(t *testing.T) {

}

func TestUserRepository_GetUserById(t *testing.T) {

}

func TestUserRepository_GetUserByUsername(t *testing.T) {

}

func TestUserRepository_LoginUser(t *testing.T) {

}

func TestUserRepository_UpdateUser(t *testing.T) {

}
