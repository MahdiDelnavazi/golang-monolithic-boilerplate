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

var repo *UserRepository

func init() {
	Test.OpenTestingDatabase()
	repo = NewUserRepository()
}

func TestUserRepository_ChangeActiveStatus(t *testing.T) {
	require.NotNil(t, repo)

	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	hashedPassword, err := Helper.HashPassword(password)
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)
	require.True(t, user.IsActive)

	userId := user.ID.Hex()
	changedStatusUser, changeStatusErr := repo.ChangeActiveStatus(userId)
	require.NoError(t, changeStatusErr)
	require.NotEmpty(t, changedStatusUser)
	require.Equal(t, changedStatusUser.UserName, user.UserName)
	require.False(t, changedStatusUser.IsActive)

}

func TestUserRepository_ChangePassword(t *testing.T) {
	require.NotNil(t, repo)

	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we change account password
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	rand.Seed(time.Now().UnixNano())
	userId := user.ID.Hex()
	ChangePassword := Request.ChangePasswordRequest{Id: userId, CurrentPassword: password, NewPassword: Helper.RandomString(8)}
	require.NotNil(t, repo)

	userChangedPass, err := repo.ChangePassword(ChangePassword)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, user.UserName, userChangedPass.UserName)
	require.Equal(t, user.UserName, userChangedPass.UserName)
	require.NotEqual(t, user.Password, ChangePassword.NewPassword)

}

func TestUserRepository_CreateUser(t *testing.T) {
	require.NotNil(t, repo)

	rand.Seed(time.Now().UnixNano())
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}

	user, err := repo.CreateUser(creatUserRequest)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)
}

func TestUserRepository_CheckUserName(t *testing.T) {
	require.NotNil(t, repo)

	rand.Seed(time.Now().UnixNano())
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	repo := NewUserRepository()

	_, err := repo.CheckUserName(creatUserRequest)
	require.NoError(t, err)
}

func TestUserRepository_GetAllUsers(t *testing.T) {
	require.NotNil(t, repo)

	rand.Seed(time.Now().UnixNano())
	getUsersReq := Request.GetAllUsers{Limit: 10, Page: 1}

	_, err := repo.GetAllUsers(getUsersReq.Limit, getUsersReq.Page)
	require.NoError(t, err)
}

func TestUserRepository_GetUserById(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we read it from db
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	getUser, err := repo.GetUserById(user.ID.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, getUser)
	require.Equal(t, user.UserName, getUser.UserName)
	require.Equal(t, user.IsActive, getUser.IsActive)
}

func TestUserRepository_GetUserByUsername(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we read it from db
	creatUserRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	getUser, err := repo.GetUserByUsername(user.UserName)
	require.NoError(t, err)
	require.NotEmpty(t, getUser)
	require.Equal(t, user.UserName, getUser.UserName)
	require.Equal(t, user.IsActive, getUser.IsActive)
}

func TestUserRepository_LoginUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	username := Helper.RandomString(5)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we read it from db and test login
	creatUserRequest := Request.CreateUserRequest{UserName: username, Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	loginRequest := Request.LoginUserRequest{UserName: username, Password: password}
	login, loginErr := repo.LoginUser(loginRequest)
	require.NoError(t, loginErr)
	require.NotEmpty(t, login)
	require.NotNil(t, login)
	require.Equal(t, login.UserName, user.UserName)

}

func TestUserRepository_UpdateUser(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	password := Helper.RandomString(8)
	username := Helper.RandomString(5)
	hashedPassword, err := Helper.HashPassword(password)
	require.Nil(t, err)
	require.NoError(t, err)

	// first we create new user then we read it from db and test login
	creatUserRequest := Request.CreateUserRequest{UserName: username, Password: hashedPassword}

	user, err := repo.CreateUser(creatUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, creatUserRequest.UserName, user.UserName)

	updateRequest := Request.UpdateUserRequest{ID: user.ID.Hex(), UserName: Helper.RandomString(5)}
	updatedUser, updateErr := repo.UpdateUser(updateRequest)
	require.NoError(t, updateErr)
	require.NotNil(t, updatedUser)
	require.NotEmpty(t, updatedUser)
	require.NotEqual(t, user.UserName, updatedUser.UserName)

}
