package Permission

import (
	"github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var permissionRepository *PermissionRepository

func init() {
	Test.OpenTestingDatabase()
	permissionRepository = NewPermissionRepository()
}

func TestPermissionRepository_CreateNewPermission(t *testing.T) {
	createPermissionRequest := Request.CreatePermission{Name: userCreate}
	permission, err := permissionRepository.CreateNewPermission(createPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, permission)
	require.NotNil(t, permission)
}

func TestPermissionRepository_GetPermissions(t *testing.T) {
	createPermissionRequest := Request.CreatePermission{Name: userCreate}
	permission, err := permissionRepository.CreateNewPermission(createPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, permission)
	require.NotNil(t, permission)

	getPermissionRequest := Request.GetAllPermissions{Limit: 1, Page: 1}
	getPermission, err := permissionRepository.GetPermissions(getPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, getPermission)
	require.NotNil(t, getPermission)
}
