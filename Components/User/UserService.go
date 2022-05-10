package Controller

import (
	"errors"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/User/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"github.com/mahdidl/golang_boilerplate/Components/User/Response"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (userService *UserService) Create(createUserRequest Request.CreateUserRequest) (Entity.User, error) {
	// check if user exist return error .
	checkUserName, _ := userService.userRepository.CheckUserName(createUserRequest)
	if checkUserName.UserName != "" {
		return Entity.User{UserName: checkUserName.UserName}, errors.New("user exist")
	}

	password, err := Helper.HashPassword(createUserRequest.Password)
	if err != nil {
		return Entity.User{}, err
	}
	createUserRequest.Password = password
	userResponse, userRepositoryError := userService.userRepository.CreateUser(createUserRequest)
	if userRepositoryError != nil {
		return Entity.User{}, userRepositoryError
	}

	return userResponse, nil
}

func (userService *UserService) GetUser(getUserRequest Request.GetUserRequest) (Response.GetUserResponse, error) {
	user, getUserError := userService.userRepository.GetUserByUsername(getUserRequest.UserName)
	if getUserError != nil {
		return Response.GetUserResponse{}, getUserError
	}
	// we need a transformer
	return Response.GetUserResponse{UserId: user.ID, UserName: user.UserName}, nil
}

func (userService *UserService) GetUserById(getUserRequest string) (Entity.User, error) {
	user, getUserError := userService.userRepository.GetUserById(getUserRequest)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService *UserService) UpdateUser(request Request.UpdateUserRequest, userId string) (Entity.User, error) {

	user, getUserError := userService.userRepository.UpdateUser(request, userId)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService *UserService) ChangePassword(request Request.ChangePasswordRequest, userId string) (Entity.User, error) {
	user, getUserError := userService.userRepository.ChangePassword(request, userId)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService *UserService) ChangeActiveStatus(userId string) (Entity.User, error) {
	user, getUserError := userService.userRepository.ChangeActiveStatus(userId)
	if getUserError != nil {
		return Entity.User{}, getUserError
	}

	return user, nil
}

func (userService *UserService) GetAllUsers(page int, limit int) (Response.ResponseAllUsers, error) {

	listUsers, err := userService.userRepository.GetAllUsers(page, limit)
	if err != nil {
		return Response.ResponseAllUsers{}, err
	}

	return Response.ResponseAllUsers{Users: listUsers}, nil
}
