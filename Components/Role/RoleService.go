package Role

import (
	"github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"github.com/mahdidl/golang_boilerplate/Components/Role/Response"
)

type RoleService struct {
	roleRepository *RoleRepository
}

func NewRoleService(permissionRepository *RoleRepository) *RoleService {
	return &RoleService{roleRepository: permissionRepository}
}

func (roleService *RoleService) Create(request Request.CreateRole) (response Entity.Role, err error) {

	role, err := roleService.roleRepository.Create(request)
	if err != nil {
		return Entity.Role{}, err
	}

	return role, nil
}

func (roleService *RoleService) GetAll(request Request.GetAllRole) (response Response.GetAllRoles, err error) {
	roles, err := roleService.roleRepository.Get(request)
	if err != nil {
		return Response.GetAllRoles{}, err
	}

	return Response.GetAllRoles{Roles: roles}, nil
}

func (roleService *RoleService) GetRoleById(Id string) (Response.GetRole, error) {
	roles, err := roleService.roleRepository.GetRoleById(Id)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: roles}, nil
}

func (roleService *RoleService) Update(request Request.UpdateRole, roleId string) (Response.GetRole, error) {
	role, err := roleService.roleRepository.Update(request, roleId)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: role}, nil
}

func (roleService *RoleService) Delete(roleId string) (Response.GetRole, error) {
	role, err := roleService.roleRepository.Delete(roleId)
	if err != nil {
		return Response.GetRole{}, err
	}

	return Response.GetRole{Roles: role}, nil
}
