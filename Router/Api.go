package Router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang_monolithic_bilerplate/Common/Config"
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

	newAuthController := Controller3.NewAuthController()

	authRoutes := routerTicket.Group("/").Use(Middleware.AuthMiddleware())
	fmt.Println("befor run user end point : ", Config.PostgresDB)
	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newUserController.LoginUser)
	routerUser.POST("/logout", newAuthUserController.Logout)
	authUser.POST("/newToken", newAuthController.AccessToken)
	authRoutes.POST("/create", newTicketController.CreateTicket)

}
