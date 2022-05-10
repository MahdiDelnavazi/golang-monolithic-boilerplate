package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	Request "github.com/mahdidl/golang_boilerplate/Components/AuthUser/Request"

	"log"

	"net/http"
)

type AuthUserController struct {
	authUserService *AuthUserService
}

func NewAuthUserController(authUserService *AuthUserService) *AuthUserController {
	return &AuthUserController{authUserService: authUserService}
}

// LogoutUser
// @Summary      Logout user
// @Description  Logout user with access token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        LogoutUserRequest  body      Request.LogoutRequest  true  "logout user"
// @Success      200                {object}  Response.GeneralResponse{data=string}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when access token is not valid"
// @Router       /auth/logout [post]
//
// Logout user with access token
func (authUserController *AuthUserController) Logout(context *gin.Context) {
	var userRequest Request.LogoutRequest
	Helper.Decode(context.Request, &userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	logoutResponse, logoutResponseError := authUserController.authUserService.LogoutUser(userRequest)

	if logoutResponseError != nil {
		response := Response.GeneralResponse{Error: true, Message: logoutResponseError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: logoutResponse}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
