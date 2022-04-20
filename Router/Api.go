package Router

import (
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Middleware"
	Controller3 "golang_monolithic_bilerplate/Components/Auth"
	Controller2 "golang_monolithic_bilerplate/Components/AuthUser"
	"golang_monolithic_bilerplate/Components/Ticket"
	Controller "golang_monolithic_bilerplate/Components/User"
	"net/http"
)

const (
	prefix               = "/api/v1"
	userPrefix           = "/user"
	ticketPrefix         = "/ticket"
	authenticationPrefix = "/auth"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)
	routerTicket := app.Group(prefix + ticketPrefix)
	routerUser := app.Group(prefix + userPrefix)
	authUser := app.Group(prefix + authenticationPrefix)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := Controller.NewUserRepository()
	newUserService := Controller.NewUserService(newUserRepository)
	newUserController := Controller.NewUserController(newUserService)

	newAuthUserRepository := Controller2.NewAuthUserRepository()
	newAuthUserService := Controller2.NewAuthUserService(newAuthUserRepository)
	newAuthUserController := Controller2.NewAuthUserController(newAuthUserService)

	newTicketRepository := Ticket.NewTicketRepository()
	newTicketService := Ticket.NewTicketService(newUserService, newTicketRepository)
	newTicketController := Ticket.NewTicketController(newTicketService)

	newAuthService := Controller3.NewAuthService()
	newAuthController := Controller3.NewAuthController(newAuthService)

	authTicketRoutes := routerTicket.Group("/").Use(Middleware.AuthMiddleware())
	authUserRoutes := routerUser.Group("/").Use(Middleware.AuthMiddleware())

	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newUserController.LoginUser)
	routerUser.POST("/logout", newAuthUserController.Logout)
	authUser.POST("/newToken", newAuthController.AccessToken)
	authTicketRoutes.POST("/create", newTicketController.CreateTicket)
	authUserRoutes.POST("/getAll", newUserController.GetAllUsers)

}
