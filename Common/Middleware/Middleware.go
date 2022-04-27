package Middleware

import (
	"errors"
	"fmt"
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
)

// AuthMiddleware creates a gin middleware for authorization
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		fmt.Println("this is token ------------>", authorizationHeader)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		accessToken := fields[1]
		payload, err := token.MakerPaseto.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		val, _ := Config.Redis.Get(payload.Username).Result()
		if val == accessToken {
			err := fmt.Errorf("token is expired")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"errors": err})
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
