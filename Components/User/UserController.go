package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	UserResponse "github.com/mahdidl/golang_boilerplate/Components/User/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	"net/http"
)

type UserController struct {
	userService *UserService
}

func NewUserController(userService *UserService) *UserController {
	return &UserController{userService: userService}
}

// @Summary      Create user
// @Description  Create user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        CreateUserRequest  body      Request.CreateUserRequest  true  "Create user request"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user exist or password < 8 character"
// @Router       /user [post]
//
// CreateUser is a handler function which is creating user
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
	response := Response.GeneralResponse{Error: false, Message: "user have been created", Data: UserResponse.CreateUserResponse{UserName: userResponse.UserName}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// GetAllUsers
// @Summary      Get all users
// @Description  Get all users return all users with pagination
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        GetAllUserRequest  query      Request.GetAllUsers  true  "get all users request"
// @Success      200                {object}  Response.GeneralResponse{data=UserResponse.ResponseAllUsers}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or password is incorrect"
// @Failure      401                {object}  Response.GeneralResponse{data=object} "unauthorized"
// @Router       /user/ [get]
// @Security ApiKeyAuth
//
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

// GetUserById
// @Summary      Get user
// @Description  Get user return user with id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId  path      string  true  "user id"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or id is incorrect"
// @Failure      401                {object}  Response.GeneralResponse{data=object} "unauthorized"
// @Router       /user/{userId} [get]
// @Security ApiKeyAuth
//
// GetUserById return user with id
func (userController *UserController) GetUserById(context *gin.Context) {
	userId := context.Param("userId")

	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.GetUserById(userId)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// UpdateUser
// @Summary      Update user
// @Description  Update user change user fields and return user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        updateUser  body      Request.UpdateUserRequest  true  "update user request"
// @Param        userId  path      string  true  "user id"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or id is incorrect"
// @Failure      401                {object}  Response.GeneralResponse{data=object} "unauthorized"
// @Router       /user/{userId} [put]
// @Security ApiKeyAuth
//
// UpdateUser for update user with any params that client sends
func (userController *UserController) UpdateUser(context *gin.Context) {
	var userRequest Request.UpdateUserRequest
	context.ShouldBindJSON(&userRequest)

	userId := context.Param("userId")
	fmt.Println("user", userRequest.UserName, userId)
	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.UpdateUser(userRequest, userId)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// ChangeUserActiveStatus
// @Summary      Change user active status
// @Description   Change user active status with id
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId  path      string  true  "user id"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or id is incorrect"
// @Failure      401                {object}  Response.GeneralResponse{data=object} "unauthorized"
// @Router       /user/{userId} [patch]
// @Security ApiKeyAuth
//
// ChangeActiveStatus for changing active and deactivate user
func (userController *UserController) ChangeActiveStatus(context *gin.Context) {
	userId := context.Param("userId")

	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.ChangeActiveStatus(userId)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// ChangeUserPassword
// @Summary      Change user password
// @Description   Change user password
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId  path      string  true  "user id"
// @Param        password  body      Request.ChangePasswordRequest  true  "change user password request"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or id is incorrect or password in incorrect"
// @Failure      401                {object}  Response.GeneralResponse{data=object} "unauthorized"
// @Router       /user/{userId}/password [put]
// @Security ApiKeyAuth
//
// ChangeActiveStatus for changing active and deactivate user
func (userController *UserController) ChangePassword(context *gin.Context) {
	var userRequest Request.ChangePasswordRequest
	//Helper.Decode(context.Request, &userRequest)
	context.ShouldBindJSON(&userRequest)

	userId := context.Param("userId")
	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	validationError := Validator.ValidationCheck(userRequest)

	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	result, responseError := userController.userService.ChangePassword(userRequest, userId)
	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: result}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
