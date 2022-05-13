package UserRole

import (
	"github.com/mahdidl/golang_boilerplate/Components/UserRole/Entity"
)

type UserRoleService struct {
	userRoleRepository *UserRoleRepository
}

func NewUserRoleService(userRoleRepository *UserRoleRepository) *UserRoleService {
	return &UserRoleService{userRoleRepository: userRoleRepository}
}

func (userRoleService UserRoleService) AttachRole(userId string, roleId string) (Entity.UserRole, error) {
	role, err := userRoleService.userRoleRepository.Attach(userId, roleId)
	if err != nil {
		return Entity.UserRole{}, err
	}

	return role, nil
}

func (userRoleService UserRoleService) DetachRole(userId string) (Entity.UserRole, error) {
	role, err := userRoleService.userRoleRepository.Detach(userId)
	if err != nil {
		return Entity.UserRole{}, err
	}

	return role, nil
}
