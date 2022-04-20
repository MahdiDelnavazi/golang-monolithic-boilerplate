package Controller

import (
	token "golang_monolithic_bilerplate/Common/Token"
	User "golang_monolithic_bilerplate/Components/AuthUser/Request"
)

type AuthUserService struct {
	AuthUserRepository *AuthUserRepository
}

func NewAuthUserService(authUserRepository *AuthUserRepository) *AuthUserService {
	return &AuthUserService{}
}

func (authUserService AuthUserService) LogoutUser(request User.LogoutRequest) (response string, err error) {
	payload, _ := token.MakerPaseto.VerifyToken(request.Token)

	err = authUserService.AuthUserRepository.LogOut(request, payload)
	if err != nil {
		return "logout failed", err
	}

	return "logout successfully", err
}
