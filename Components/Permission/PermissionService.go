package Permission

import (
	Request "github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	Response "github.com/mahdidl/golang_boilerplate/Components/Permission/Response"
)

type PermissionService struct {
	permissionRepository *PermissionRepository
}

func NewPermissionService(permissionRepository *PermissionRepository) *PermissionService {
	return &PermissionService{permissionRepository: permissionRepository}
}

func (permissionService PermissionService) CreateNewPermission(request Request.CreatePermission) (response Response.CreatePermission, err error) {

	permission, err := permissionService.permissionRepository.CreateNewPermission(request)
	if err != nil {
		return Response.CreatePermission{}, err
	}

	return Response.CreatePermission{ID: permission.Id, Name: permission.Name}, nil
}

func (permissionService PermissionService) GetPermissions(request Request.GetAllPermissions) (response Response.GetPermissions, err error) {
	permissions, err := permissionService.permissionRepository.GetPermissions(request)
	if err != nil {
		return Response.GetPermissions{}, err
	}

	return Response.GetPermissions{Permissions: permissions}, nil
}
