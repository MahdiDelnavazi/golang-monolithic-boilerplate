package Permission

import (
	"github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var permissionService *PermissionService

func init() {
	Test.OpenTestingDatabase()
	permissionService = NewPermissionService(NewPermissionRepository())
}

func TestPermissionService_CreateNewPermission(t *testing.T) {
	createPermissionRequest := Request.CreatePermission{Name: userCreate}
	permission, err := permissionService.CreateNewPermission(createPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, permission)
	require.NotNil(t, permission)
}

func TestPermissionService_GetPermissions(t *testing.T) {
	createPermissionRequest := Request.CreatePermission{Name: userCreate}
	permission, err := permissionService.CreateNewPermission(createPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, permission)
	require.NotNil(t, permission)

	getPermissionRequest := Request.GetAllPermissions{Limit: 1, Page: 1}
	getPermission, err := permissionService.GetPermissions(getPermissionRequest)
	require.NoError(t, err)
	require.NotEmpty(t, getPermission)
	require.NotNil(t, getPermission)
}
