package Middleware

import (
	"errors"
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	token "github.com/mahdidl/golang_boilerplate/Common/Token"

	"github.com/mahdidl/golang_boilerplate/Common/Config"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
	isDevelop               = true
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !isDevelop {
			authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

			if len(authorizationHeader) == 0 {
				err := errors.New("authorization header is not provided")
				response := Response.GeneralResponse{Error: true, Message: err.Error()}
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": response})
				return
			}

			fields := strings.Fields(authorizationHeader)
			if len(fields) < 2 {
				err := errors.New("invalid authorization header format")
				response := Response.GeneralResponse{Error: true, Message: err.Error()}
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": response})
				return
			}

			authorizationType := strings.ToLower(fields[0])
			if authorizationType != authorizationTypeBearer {
				err := fmt.Errorf("unsupported authorization type %s", authorizationType)
				response := Response.GeneralResponse{Error: true, Message: err.Error()}
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": response})
				return
			}

			accessToken := fields[1]
			payload, err := token.MakerPaseto.VerifyToken(accessToken)
			if err != nil {
				response := Response.GeneralResponse{Error: true, Message: err.Error()}
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": response})
				return
			}

			val, _ := Config.Redis.Get(payload.Username).Result()
			if val == accessToken {
				err := fmt.Errorf("token is expired")
				response := Response.GeneralResponse{Error: true, Message: err.Error()}
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"response": response})
				return
			}

			ctx.Set(authorizationPayloadKey, payload)
			ctx.Next()
		}
	}
}
