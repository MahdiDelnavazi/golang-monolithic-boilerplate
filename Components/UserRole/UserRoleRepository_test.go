package UserRole

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Role"
	RequestRole "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	User "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	RequestAttach "github.com/mahdidl/golang_boilerplate/Components/UserRole/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var userRoleRepository *UserRoleRepository
var userRepository *User.UserRepository
var roleRepository *Role.RoleRepository

func init() {
	Test.OpenTestingDatabase()
	userRoleRepository = NewUserRoleRepository()
	roleRepository = Role.NewRoleRepository()
	userRepository = User.NewUserRepository()
}

func TestUserRoleRepository_Attach(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	user, err := userRepository.CreateUser(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotNil(t, user)
	require.Equal(t, user.UserName, userRequest.UserName)
	require.Equal(t, user.RoleID, primitive.NilObjectID)

	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleRepository.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	attachRoleToUserRequest := RequestAttach.AttachRole{UserId: user.ID.Hex(), RoleId: role.Id.Hex()}
	userAttachedRole, err := userRoleRepository.Attach(attachRoleToUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, userAttachedRole)
	require.NotNil(t, userAttachedRole)
	require.NotEqual(t, userAttachedRole.RoleID, primitive.NilObjectID)
}

func TestUserRoleRepository_Detach(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	user, err := userRepository.CreateUser(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotNil(t, user)
	require.Equal(t, user.UserName, userRequest.UserName)
	require.Equal(t, user.RoleID, primitive.NilObjectID)

	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleRepository.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	attachRoleToUserRequest := RequestAttach.AttachRole{UserId: user.ID.Hex(), RoleId: role.Id.Hex()}
	userAttachedRole, err := userRoleRepository.Attach(attachRoleToUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, userAttachedRole)
	require.NotNil(t, userAttachedRole)
	require.NotEqual(t, userAttachedRole.RoleID, primitive.NilObjectID)

	detachRoleToUserRequest := RequestAttach.DetachRole{UserId: user.ID.Hex()}
	userDetachedRole, err := userRoleService.DetachRole(detachRoleToUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, userDetachedRole)
	require.NotNil(t, userDetachedRole)
	require.Equal(t, userDetachedRole.RoleID, primitive.NilObjectID)
}
