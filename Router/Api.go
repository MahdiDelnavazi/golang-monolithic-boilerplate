package Router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	prefix = "/api/v1"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//newUserRepository := Repository.NewUserRepository(log, db)

	//authRoutes := router.Group("/").Use(Middleware.AuthMiddleware(token, redis))
	//router.POST("/create", newUserController.CreateUser)

}
