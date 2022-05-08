package RolePermission

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Permission"
	"github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"github.com/mahdidl/golang_boilerplate/Components/Role"
	RequestRole "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	Request2 "github.com/mahdidl/golang_boilerplate/Components/RolePermission/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var rolePermissionService *RolePermissionService
var roleService *Role.RoleService
var permissionService *Permission.PermissionService

func init() {
	Test.OpenTestingDatabase()
	rolePermissionService = NewRolePermissionService(NewRolePermissionRepository())
	roleService = Role.NewRoleService(Role.NewRoleRepository())
	permissionService = Permission.NewPermissionService(Permission.NewPermissionRepository())
}

func TestRolePermissionService_Attach(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)
	require.Empty(t, role.PermissionsId)

	permissionName := Helper.RandomString(5)
	createPermissionRequest := Request.CreatePermission{Name: permissionName}
	permission, err := permissionService.CreateNewPermission(createPermissionRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, permission.Name, permissionName)

	attachRolePermissionRequest := Request2.AttachPermission{PermissionId: permission.ID.Hex(), RoleId: role.Id.Hex()}
	AttachedRolePermission, err := rolePermissionService.Attach(attachRolePermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.NotEqual(t, AttachedRolePermission.PermissionsId[0], primitive.NilObjectID)

}

func TestRolePermissionService_Detach(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)
	require.Empty(t, role.PermissionsId)

	permissionName := Helper.RandomString(5)
	createPermissionRequest := Request.CreatePermission{Name: permissionName}
	permission, err := permissionService.CreateNewPermission(createPermissionRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, permission.Name, permissionName)

	attachRolePermissionRequest := Request2.AttachPermission{PermissionId: permission.ID.Hex(), RoleId: role.Id.Hex()}
	AttachedRolePermission, err := rolePermissionService.Attach(attachRolePermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.NotEqual(t, AttachedRolePermission.PermissionsId[0], primitive.NilObjectID)

	detachRolePermissionRequest := Request2.DetachPermission{PermissionId: permission.ID.Hex(), RoleId: role.Id.Hex()}
	detachedRolePermission, err := rolePermissionService.Detach(detachRolePermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Empty(t, detachedRolePermission.PermissionsId)
}
