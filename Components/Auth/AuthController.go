package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	AuthRequest "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	AuthResponse "github.com/mahdidl/golang_boilerplate/Components/Auth/Response"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	UserResponse "github.com/mahdidl/golang_boilerplate/Components/User/Response"
	"log"

	"net/http"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}

// @Summary      New access token
// @Description  New access token with refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        AccessTokenRequest  body      AuthRequest.AccessTokenRequest  true  "for get new access token"
// @Success      200                {object}  Response.GeneralResponse{data=User.AccessTokenRequest}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /authentication/newToken [post]
//
// LoginUser for get access token
func (authController *AuthController) AccessToken(context *gin.Context) {
	var accessTokenReq AuthRequest.AccessTokenRequest
	Helper.Decode(context.Request, &accessTokenReq)

	validationError := Validator.ValidationCheck(accessTokenReq)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	token, err := authController.authService.CreateAccessToken(accessTokenReq)
	if err != nil {
		response := Response.GeneralResponse{Error: true, Message: err.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: "successful", Data: AuthResponse.AccessTokenResponse{AccessToken: token}}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}

// LoginUser
// @Summary      Login user
// @Description  Login user with username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        LoginUserRequest  body      Request.CreateUserRequest  true  "Create user request"
// @Success      200                {object}  Response.GeneralResponse{data=UserResponse.LoginUserResponse}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "when user not exist or password is incorrect"
// @Router       /authentication/login [post]
//
// LoginUser for get access token
func (authController *AuthController) LoginUser(context *gin.Context) {
	var userRequest Request.LoginUserRequest
	context.ShouldBindJSON(&userRequest)

	validationError := Validator.ValidationCheck(userRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userResponse, responseError := authController.authService.LoginUser(userRequest)

	if responseError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"response": Response.ErrorResponse{Error: responseError.Error()}})
		return
	}

	// all ok
	// create general response
	var loginResponse UserResponse.LoginUserResponse
	loginResponse = userResponse
	response := Response.GeneralResponse{Error: false, Message: "your login is successful", Data: loginResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
