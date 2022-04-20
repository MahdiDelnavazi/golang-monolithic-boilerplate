package Controller

import (
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Helper"
	Response "golang_monolithic_bilerplate/Common/Response"
	"golang_monolithic_bilerplate/Common/Validator"
	"golang_monolithic_bilerplate/Components/User/Request"
	UserResponse "golang_monolithic_bilerplate/Components/User/Response"
	"log"
	"net/http"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

func (userControler *UserController) CreateUser(context *gin.Context) {
	var userRequest Request.CreateUserRequest
	Helper.Decode(context.Request, &userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
	}

	userResponse, responseError := userControler.userService.Create(userRequest)

	if responseError != nil {
		// if username not empty means its userExist error
		if userResponse.UserName != "" {
			response := Response.GeneralResponse{Error: true, Message: "user exist", Data: nil}
			context.JSON(http.StatusBadRequest, gin.H{"response": response})
			return
		}
		// if username is empty means its validation error
		context.JSON(http.StatusBadRequest, gin.H{"response": responseError})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "user have been created", Data: UserResponse.CreateUserResponse{UserName: userResponse.UserName}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (userControler *UserController) LoginUser(context *gin.Context) {
	var userRequest Request.LoginUserRequest
	Helper.Decode(context.Request, &userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
	}

	userResponse, responseError := userControler.userService.LoginUser(userRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": responseError})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "your login is successful", Data: UserResponse.LoginUserResponse{UserName: userResponse.UserName, AccessToken: userResponse.AccessToken, RefreshToken: userResponse.RefreshToken}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

//func (userControler *UserController) Logout(context *gin.Context) {
//	var userRequest User.LogoutRequest
//	Helper.Decode(context.Request, &userRequest)
//
//	logoutResponse, logoutResponseError := userControler.userService.LogoutUser(userRequest)
//
//	if logoutResponseError != nil {
//		context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
//		return
//	}
//
//	// all ok
//	// create general response
//	response1 := Response.GeneralResponse{Error: false, Message: logoutResponse}
//	context.JSON(http.StatusOK, gin.H{"response": response1})
//}
