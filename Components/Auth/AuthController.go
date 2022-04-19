package Controller

import (
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Helper"
	"golang_monolithic_bilerplate/Common/Response"
	token "golang_monolithic_bilerplate/Common/Token"
	User "golang_monolithic_bilerplate/Components/Auth/Request"
	Response2 "golang_monolithic_bilerplate/Components/Auth/Response"

	"net/http"
	"time"
)

type AuthController struct {

	//userService *Service.UserService
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (authController *AuthController) AccessToken(context *gin.Context) {
	var accessTokenReq User.AccessTokenRequest
	Helper.Decode(context.Request, &accessTokenReq)

	//logoutResponse, logoutResponseError := userControler.userService.LogoutUser(userRequest)

	//if logoutResponseError != nil {
	//	context.JSON(http.StatusBadRequest, gin.H{"response": logoutResponseError})
	//	return
	//}

	payload, err := token.MakerPaseto.VerifyToken(accessTokenReq.AccessToken)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}
	newToken, err := token.MakerPaseto.CreateToken(payload.Username, time.Hour)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": err})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: "successful", Data: Response2.AccessTokenResponse{AccessToken: newToken}}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
