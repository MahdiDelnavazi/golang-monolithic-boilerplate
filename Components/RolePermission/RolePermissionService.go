package RolePermission

import (
	Entity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
)

type RolePermissionService struct {
	RolePermissionRepository *RolePermissionRepository
}

func NewRolePermissionService(rolePermissionRepository *RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{RolePermissionRepository: rolePermissionRepository}
}

func (rolePermissionService RolePermissionService) Attach(roleId string, permissionId string) (Entity.Role, error) {
	role, err := rolePermissionService.RolePermissionRepository.Attach(roleId, permissionId)
	if err != nil {
		return Entity.Role{}, err
	}

	return role, nil
}

func (rolePermissionService RolePermissionService) Detach(roleId string, permissionId string) (Entity.Role, error) {
	role, err := rolePermissionService.RolePermissionRepository.Detach(roleId, permissionId)
	if err != nil {
		return Entity.Role{}, err
	}

	return role, nil
}
