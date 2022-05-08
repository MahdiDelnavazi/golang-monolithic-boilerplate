package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	Response3 "github.com/mahdidl/golang_boilerplate/Components/User/Response"

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

	context.ShouldBindJSON(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
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
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "user have been created", Data: Response3.CreateUserResponse{UserName: userResponse.UserName}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// LoginUser for get access token
func (userControler *UserController) LoginUser(context *gin.Context) {
	var userRequest Request.LoginUserRequest
	context.ShouldBindJSON(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userResponse, responseError := userControler.userService.LoginUser(userRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "your login is successful", Data: userResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// GetAllUsers return all users with pagination
func (userController *UserController) GetAllUsers(context *gin.Context) {
	var userRequest Request.GetAllUsers
	context.ShouldBindQuery(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.GetAllUsers(userRequest.Page, userRequest.Limit)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": responseError})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// GetUserById return user with id
func (userController *UserController) GetUserById(context *gin.Context) {
	var userRequest Request.GetUser

	context.ShouldBindQuery(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.GetUserById(userRequest)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// UpdateUser for update user with any params that client sends
func (userController *UserController) UpdateUser(context *gin.Context) {
	var userRequest Request.UpdateUserRequest
	context.ShouldBindJSON(&userRequest)
	context.ShouldBindQuery(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.UpdateUser(userRequest)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// ChangeActiveStatus for changing active and deactivate user
func (userController *UserController) ChangeActiveStatus(context *gin.Context) {
	var userRequest Request.ChangeStatusRequest
	context.ShouldBindQuery(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.ChangeActiveStatus(userRequest)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (userController *UserController) ChangePassword(context *gin.Context) {
	var userRequest Request.ChangePasswordRequest
	//Helper.Decode(context.Request, &userRequest)
	context.ShouldBindJSON(&userRequest)
	context.ShouldBindQuery(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.ChangePassword(userRequest)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
