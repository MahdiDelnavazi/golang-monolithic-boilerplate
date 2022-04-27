package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	User "github.com/mahdidl/golang_boilerplate/Components/AuthUser/Request"

	"log"

	"net/http"
)

type AuthUserController struct {
	authUserService *AuthUserService
}

func NewAuthUserController(authUserService *AuthUserService) *AuthUserController {
	return &AuthUserController{authUserService: authUserService}
}

func (authUserController *AuthUserController) Logout(context *gin.Context) {
	var userRequest User.LogoutRequest
	Helper.Decode(context.Request, &userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
	}

	logoutResponse, logoutResponseError := authUserController.authUserService.LogoutUser(userRequest)

	if logoutResponseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: logoutResponse}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
