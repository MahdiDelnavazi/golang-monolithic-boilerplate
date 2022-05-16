package RolePermission

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Permission"
	"github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"github.com/mahdidl/golang_boilerplate/Components/Role"
	RequestRole "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

var rolePermissionRepository *RolePermissionRepository
var roleRepository *Role.RoleRepository
var permissionRepository *Permission.PermissionRepository

func init() {
	Test.OpenTestingDatabase()
	rolePermissionRepository = NewRolePermissionRepository()
	roleRepository = Role.NewRoleRepository()
	permissionRepository = Permission.NewPermissionRepository()
}

func TestRolePermissionRepository_Attach(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleRepository.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)
	require.Empty(t, role.PermissionsId)

	permissionName := Helper.RandomString(5)
	createPermissionRequest := Request.CreatePermission{Name: permissionName}
	permission, err := permissionRepository.CreateNewPermission(createPermissionRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, permission.Name, permissionName)

	AttachedRolePermission, err := rolePermissionRepository.Attach(permission.Id.Hex(), role.Id.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.NotEqual(t, AttachedRolePermission.PermissionsId[0], primitive.NilObjectID)
}

func TestRolePermissionRepository_Detach(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := RequestRole.CreateRole{Name: roleName}
	role, err := roleRepository.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)
	require.Empty(t, role.PermissionsId)

	permissionName := Helper.RandomString(5)
	createPermissionRequest := Request.CreatePermission{Name: permissionName}
	permission, err := permissionRepository.CreateNewPermission(createPermissionRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, permission.Name, permissionName)

	AttachedRolePermission, err := rolePermissionRepository.Attach(permission.Id.Hex(), role.Id.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.NotEqual(t, AttachedRolePermission.PermissionsId[0], primitive.NilObjectID)

	detachedRolePermission, err := rolePermissionRepository.Detach(permission.Id.Hex(), role.Id.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Empty(t, detachedRolePermission.PermissionsId)
}
