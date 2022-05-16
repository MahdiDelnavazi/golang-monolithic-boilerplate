package Role

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"math/rand"
	"testing"
	"time"
)

var roleRepository *RoleRepository

func init() {
	Test.OpenTestingDatabase()
	roleRepository = NewRoleRepository()
}

func TestRoleRepository_Create(t *testing.T) {
	require.NotNil(t, roleRepository)

	roleName := Helper.RandomString(5)

	roleReq := Request.CreateRole{Name: roleName}

	role, err := roleRepository.Create(roleReq)
	require.NoError(t, err)
	require.NotNil(t, role)
	require.NotEmpty(t, role)
	require.Equal(t, role.Name, roleReq.Name)

}

func TestRoleRepository_Delete(t *testing.T) {
	require.NotNil(t, roleRepository)

	roleName := Helper.RandomString(5)

	roleReq := Request.CreateRole{Name: roleName}

	role, err := roleRepository.Create(roleReq)
	require.NoError(t, err)
	require.NotNil(t, role)
	require.NotEmpty(t, role)
	require.Equal(t, role.Name, roleReq.Name)

	id := role.Id.Hex()

	roleDeleted, err := roleRepository.Delete(id)
	require.NoError(t, err)
	require.Equal(t, role.Id.Hex(), id)
	require.NotEmpty(t, roleDeleted)
	require.NotNil(t, roleDeleted.DeletedAt)

}

func TestRoleRepository_Get(t *testing.T) {
	require.NotNil(t, roleRepository)
	rand.Seed(time.Now().UnixNano())

	roleName := Helper.RandomString(5)

	roleReq := Request.CreateRole{Name: roleName}

	role, err := roleRepository.Create(roleReq)
	require.NoError(t, err)
	require.NotNil(t, role)
	require.NotEmpty(t, role)
	require.Equal(t, role.Name, roleReq.Name)

	getAllRolesReq := Request.GetAllRole{Page: 1, Limit: 10}
	_, err = roleRepository.Get(getAllRolesReq)
	require.NoError(t, err)

}

func TestRoleRepository_GetRoleById(t *testing.T) {
	require.NotNil(t, roleRepository)
	rand.Seed(time.Now().UnixNano())

	roleName := Helper.RandomString(5)

	roleReq := Request.CreateRole{Name: roleName}

	role, err := roleRepository.Create(roleReq)
	require.NoError(t, err)
	require.NotNil(t, role)
	require.NotEmpty(t, role)
	require.Equal(t, role.Name, roleReq.Name)

	id := role.Id.Hex()

	getRole, err := roleRepository.GetRoleById(id)
	require.NoError(t, err)
	require.Equal(t, role.Id, getRole.Id)
	require.NotEmpty(t, getRole)
}

func TestRoleRepository_Update(t *testing.T) {
	require.NotNil(t, roleRepository)
	rand.Seed(time.Now().UnixNano())

	roleName := Helper.RandomString(5)

	roleReq := Request.CreateRole{Name: roleName}

	role, err := roleRepository.Create(roleReq)
	require.NoError(t, err)
	require.NotNil(t, role)
	require.NotEmpty(t, role)
	require.Equal(t, role.Name, roleReq.Name)

	id := role.Id.Hex()
	rand.Seed(time.Now().UnixNano())
	updateReq := Request.UpdateRole{Name: Helper.RandomString(5)}

	updatedRole, err := roleRepository.Update(updateReq, id)
	require.NoError(t, err)
	require.Equal(t, role.Id, updatedRole.Id)
	require.NotEmpty(t, updatedRole)
}
