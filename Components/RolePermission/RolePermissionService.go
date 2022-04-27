package RolePermission

import (
	Entity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/RolePermission/Request"
)

type RolePermissionService struct {
	RolePermissionRepository *RolePermissionRepository
}

func NewRolePermissionService(rolePermissionRepository *RolePermissionRepository) *RolePermissionService {
	return &RolePermissionService{RolePermissionRepository: rolePermissionRepository}
}

func (rolePermissionService RolePermissionService) Attach(request Request.AttachPermission) (Entity.Role, error) {
	role, err := rolePermissionService.RolePermissionRepository.Attach(request)
	if err != nil {
		return Entity.Role{}, err
	}

	return role, nil
}
