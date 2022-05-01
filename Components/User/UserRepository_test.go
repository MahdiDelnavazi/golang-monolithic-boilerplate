package Controller

import (
	"github.com/mahdidl/golang_boilerplate/Components/User/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/mahdidl/golang_boilerplate/Test"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"testing"
	"time"
)

func TestUserRepository_ChangeActiveStatus(t *testing.T) {

	Test.OpenTestingDatabase()

	var user Entity.User

	creatUserRequest := Request.CreateUserRequest{UserName: "mimdl", Password: "salamalekom"}

	result, err := Test.UserCollection.InsertOne(Test.DBCtx, Entity.User{ID: primitive.NewObjectID(), IsActive: true,
		UserName: creatUserRequest.UserName, Password: creatUserRequest.Password, CreatedAt: time.Now()})

	require.NoError(t, err)
	require.NotEmpty(t, result.InsertedID)

	err = Test.UserCollection.FindOne(Test.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&user)
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
