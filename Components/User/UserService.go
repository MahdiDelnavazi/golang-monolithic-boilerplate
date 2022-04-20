package Controller

import (
	"errors"
	"fmt"
	token "golang_monolithic_bilerplate/Common/Token"
	"golang_monolithic_bilerplate/Common/Validator"
	"golang_monolithic_bilerplate/Components/User/Request"
	"golang_monolithic_bilerplate/Components/User/Response"

	"log"
	"time"
)

type UserService struct {
	userRepository *UserRepository
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{}
}

func (userService UserService) Create(createUserRequest Request.CreateUserRequest) (Response.CreateUserResponse, error) {

	// validate username len and not empty
	validationError := Validator.ValidationCheck(createUserRequest)
	log.Println(validationError)
	if validationError != nil {
		return Response.CreateUserResponse{}, validationError
	}

	// check if user exist return error .
	checkUserName, _ := userService.userRepository.CheckUserName(createUserRequest)
	if checkUserName.UserName != "" {
		return Response.CreateUserResponse{UserName: checkUserName.UserName}, errors.New("user exist")
	}

	userResponse, userRepositoryError := userService.userRepository.CreateUser(createUserRequest)
	fmt.Println("user befor create in service : ", userResponse)
	if userRepositoryError != nil {
		return Response.CreateUserResponse{}, userRepositoryError
	}
	// we need a transformer
	return Response.CreateUserResponse{UserName: userResponse.UserName}, nil
}

func (userService UserService) LoginUser(loginUserRequest Request.LoginUserRequest) (Response.LoginUserResponse, error) {
	// validate username len and not empty
	validationError := Validator.ValidationCheck(loginUserRequest)

	if validationError != nil {
		return Response.LoginUserResponse{}, validationError
	}

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
	return Response.LoginUserResponse{UserName: user.UserName, AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

// goes to user auth
//func (userService UserService) LogoutUser(request User.LogoutRequest) (response string, err error) {
//	payload, _ := userService.token.VerifyToken(request.Token)
//
//	err = userService.userRepository.LogOut(request, payload)
//	if err != nil {
//		return "logout failed", err
//	}
//
//	return "logout successfully", err
//}

func (userService UserService) GetUser(getUserRequest Request.GetUserRequest) (Response.GetUserResponse, error) {
	// validate username len and not empty
	validationError := Validator.ValidationCheck(getUserRequest)

	if validationError != nil {
		return Response.GetUserResponse{}, validationError
	}

	user, getUserError := userService.userRepository.GetUser(getUserRequest)
	if getUserError != nil {
		return Response.GetUserResponse{}, getUserError
	}
	// we need a transformer
	return Response.GetUserResponse{UserId: user.ID, UserName: user.UserName}, nil
}
