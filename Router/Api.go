package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Config"
	"golang_monolithic_bilerplate/Common/Middleware"
	Controller "golang_monolithic_bilerplate/Components/User"
	"net/http"
)

const (
	prefix       = "/api/v1"
	userPrefix   = "/user"
	ticketPrefix = "/ticket"
)

//
func Routes(app *gin.Engine) {
	router := app.Group(prefix)
	routerTicket := app.Group(prefix + ticketPrefix)
	routerUser := app.Group(prefix + userPrefix)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := Controller.NewUserRepository()
	newUserService := Controller.NewUserService(newUserRepository)
	newUserController := Controller.NewUserController(newUserService)

	authRoutes := routerTicket.Group("/").Use(Middleware.AuthMiddleware())
	fmt.Println("befor run user end point : ", Config.PostgresDB)
	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newUserController.CreateUser)
	routerUser.POST("/logout", newUserController.CreateUser)
	authRoutes.POST("/create", newUserController.CreateUser)

}
