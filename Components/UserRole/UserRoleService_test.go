package UserRole

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Role"
	RequestRole "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	RequestAttach "github.com/mahdidl/golang_boilerplate/Components/UserRole/Request"
	"go.mongodb.org/mongo-driver/bson/primitive"

	User "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var userRoleService *UserRoleService
var userService *User.UserService
var roleService *Role.RoleService

func init() {
	Test.OpenTestingDatabase()
	userRoleService = NewUserRoleService(NewUserRoleRepository())
	roleService = Role.NewRoleService(Role.NewRoleRepository())
	userService = User.NewUserService(User.NewUserRepository())
}

func TestUserRoleService_AttachRole(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	user, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotNil(t, user)
	require.Equal(t, user.UserName, userRequest.UserName)
	require.Equal(t, user.RoleID, primitive.NilObjectID)

	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	attachRoleToUserRequest := RequestAttach.AttachRole{UserId: user.ID.Hex(), RoleId: role.ID.Hex()}
	userAttachedRole, err := userRoleService.AttachRole(attachRoleToUserRequest)
	require.NoError(t, err)
	require.NotEmpty(t, userAttachedRole)
	require.NotNil(t, userAttachedRole)
	require.NotEqual(t, userAttachedRole.RoleID, primitive.NilObjectID)
}

func TestUserRoleService_DetachRole(t *testing.T) {
	userRequest := Request.CreateUserRequest{UserName: Helper.RandomString(5), Password: Helper.RandomString(8)}
	user, err := userService.Create(userRequest)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotNil(t, user)
	require.Equal(t, user.UserName, userRequest.UserName)

	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	attachRoleToUserRequest := RequestAttach.AttachRole{UserId: user.ID.Hex(), RoleId: role.ID.Hex()}
	userAttachedRole, err := userRoleService.AttachRole(attachRoleToUserRequest)
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
