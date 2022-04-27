package Role

import (
	"github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"github.com/mahdidl/golang_boilerplate/Components/Role/Response"
)

type RoleService struct {
	roleRepository *RoleRepository
}

func NewPermissionService(permissionRepository *RoleRepository) *RoleService {
	return &RoleService{roleRepository: permissionRepository}
}

func (roleService RoleService) Create(request Request.CreateRole) (response Response.CreateRole, err error) {

	permission, err := roleService.roleRepository.Create(request)
	if err != nil {
		return Response.CreateRole{}, err
	}

	return Response.CreateRole{ID: permission.Id, Name: permission.Name}, nil
}

func (roleService RoleService) GetAll(request Request.GetAllRole) (response Response.GetAllRoles, err error) {
	roles, err := roleService.roleRepository.Get(request)
	if err != nil {
		return Response.GetAllRoles{}, err
	}

	return Response.GetAllRoles{Roles: roles}, nil
}

func (roleService RoleService) GetRoleById(Id string) (Response.GetRole, error) {
	roles, err := roleService.roleRepository.GetRoleById(Id)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: roles}, nil
}

func (roleService RoleService) Update(request Request.UpdateRole) (Response.GetRole, error) {
	role, err := roleService.roleRepository.Update(request)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: role}, nil
}

func (roleService RoleService) Delete(request Request.DeleteRole) (Response.GetRole, error) {
	role, err := roleService.roleRepository.Delete(request)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: role}, nil
}
