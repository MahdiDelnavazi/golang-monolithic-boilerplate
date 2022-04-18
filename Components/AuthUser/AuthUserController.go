package Controller

import (
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Helper"
	Response2 "golang_monolithic_bilerplate/Common/Response"
	User "golang_monolithic_bilerplate/Components/AuthUser/Request"

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

	logoutResponse, logoutResponseError := authUserController.authUserService.LogoutUser(userRequest)

	if logoutResponseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
		return
	}

	// all ok
	// create general response
	response1 := Response2.GeneralResponse{Error: false, Message: logoutResponse}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
