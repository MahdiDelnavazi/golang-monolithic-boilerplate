package Controller

import (
	"errors"
	"golang_monolithic_bilerplate/Common/Helper"
	token "golang_monolithic_bilerplate/Common/Token"
	"golang_monolithic_bilerplate/Components/User/Entity"
	"golang_monolithic_bilerplate/Components/User/Request"
	"golang_monolithic_bilerplate/Components/User/Response"

	"time"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService UserService) Create(createUserRequest Request.CreateUserRequest) (Response.CreateUserResponse, error) {
	// check if user exist return error .
	checkUserName, _ := userService.userRepository.CheckUserName(createUserRequest)
	if checkUserName.UserName != "" {
		return Response.CreateUserResponse{UserName: checkUserName.UserName}, errors.New("user exist")
	}

	password, err := Helper.HashPassword(createUserRequest.Password)
	if err != nil {
		return Response.CreateUserResponse{}, err
	}

	userResponse, userRepositoryError := userService.userRepository.CreateUser(createUserRequest, password)
	if userRepositoryError != nil {
		return Response.CreateUserResponse{}, userRepositoryError
	}

	return Response.CreateUserResponse{UserName: userResponse.UserName}, nil
}

func (userService UserService) LoginUser(loginUserRequest Request.LoginUserRequest) (Response.LoginUserResponse, error) {
	user, getUserError := userService.userRepository.LoginUser(loginUserRequest)
	if getUserError != nil {
		return Response.LoginUserResponse{}, getUserError
	}

	//create new token for login
	accessToken, err := token.MakerPaseto.CreateToken(loginUserRequest.UserName, time.Hour)
	if err != nil {
		return Response.LoginUserResponse{}, err
	}

	refreshToken, errRefreshToken := token.MakerPaseto.CreateToken(loginUserRequest.UserName, time.Hour*120)
	if errRefreshToken != nil {
		return Response.LoginUserResponse{}, err
	}

	// we need a transformer
	return Response.LoginUserResponse{UserName: user.UserName, AccessToken: accessToken, RefreshToken: refreshToken, ID: user.ID.Hex()}, nil
}

func (userService UserService) GetUser(getUserRequest Request.GetUserRequest) (Response.GetUserResponse, error) {
	user, getUserError := userService.userRepository.GetUserByUsername(getUserRequest.UserName)
	if getUserError != nil {
		return Response.GetUserResponse{}, getUserError
	}
	// we need a transformer
	return Response.GetUserResponse{UserId: user.ID, UserName: user.UserName}, nil
}

func (userService UserService) GetUserById(getUserRequest Request.GetUser) (Entity.User, error) {
	user, getUserError := userService.userRepository.GetUserById(getUserRequest.ID)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService UserService) UpdateUser(request Request.UpdateUserRequest) (Entity.User, error) {

	user, getUserError := userService.userRepository.UpdateUser(request)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService UserService) ChangePassword(request Request.ChangePasswordRequest) (string, error) {
	message, getUserError := userService.userRepository.ChangePassword(request)
	if getUserError != nil {
		return "", getUserError
	}

	return message, nil
}

func (userService UserService) ChangeActiveStatus(request Request.ChangeStatusRequest) (Entity.User, error) {
	user, getUserError := userService.userRepository.ChangeActiveStatus(request.ID)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService UserService) GetAllUsers(page int, limit int) (Response.ResponseAllUsers, error) {

	listUsers, err := userService.userRepository.GetAllUsers(page, limit)
	if err != nil {
		return Response.ResponseAllUsers{}, err
	}

	return Response.ResponseAllUsers{Users: listUsers}, nil
}
