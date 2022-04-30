package UserRole

import (
	"github.com/mahdidl/golang_boilerplate/Components/UserRole/Entity"
	Request "github.com/mahdidl/golang_boilerplate/Components/UserRole/Request"
)

type UserRoleService struct {
	userRoleRepository *UserRoleRepository
}

func NewUserRoleService(userRoleRepository *UserRoleRepository) *UserRoleService {
	return &UserRoleService{userRoleRepository: userRoleRepository}
}

func (userRoleService UserRoleService) AttachRole(request Request.AttachRole) (Entity.UserRole, error) {
	role, err := userRoleService.userRoleRepository.Attach(request)
	if err != nil {
		return Entity.UserRole{}, err
	}

	return role, nil
}

func (userRoleService UserRoleService) DetachRole(request Request.DetachRole) (Entity.UserRole, error) {
	role, err := userRoleService.userRoleRepository.Detach(request)
	if err != nil {
		return Entity.UserRole{}, err
	}

	return role, nil
}
