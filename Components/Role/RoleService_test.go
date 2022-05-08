package Role

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var roleService *RoleService

func init() {
	Test.OpenTestingDatabase()
	roleService = NewRoleService(NewRoleRepository())
}

func TestRoleService_Create(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := Request.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)
}

func TestRoleService_Delete(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := Request.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	deleteRoleRequest := Request.DeleteRole{Id: role.ID.Hex()}
	deleteRoleResponse, err := roleService.Delete(deleteRoleRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, deleteRoleResponse)
	require.Equal(t, role.Name, deleteRoleResponse.Roles.Name)

}

func TestRoleService_GetAll(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := Request.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	getAllRolesRequest := Request.GetAllRole{Page: 1, Limit: 1}
	_, err = roleService.GetAll(getAllRolesRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
}

func TestRoleService_GetRoleById(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := Request.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	getRole, err := roleService.GetRoleById(role.ID.Hex())
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, getRole.Roles.Id, role.ID)

}

func TestRoleService_Update(t *testing.T) {
	roleName := Helper.RandomString(5)
	createRoleRequest := Request.CreateRole{Name: roleName}
	role, err := roleService.Create(createRoleRequest)

	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.Equal(t, role.Name, roleName)

	updateRoleRequest := Request.UpdateRole{Id: role.ID.Hex(), Name: Helper.RandomString(5)}
	updatedRole, err := roleService.Update(updateRoleRequest)
	require.NoError(t, err)
	require.NotEmpty(t, role)
	require.NotNil(t, role)
	require.NotEqual(t, role.Name, updatedRole.Roles.Name)

}
