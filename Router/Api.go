package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Middleware"
	Auth "github.com/mahdidl/golang_boilerplate/Components/Auth"
	AuthUser "github.com/mahdidl/golang_boilerplate/Components/AuthUser"
	"github.com/mahdidl/golang_boilerplate/Components/Permission"
	"github.com/mahdidl/golang_boilerplate/Components/Role"
	RolePermission "github.com/mahdidl/golang_boilerplate/Components/RolePermission"
	"github.com/mahdidl/golang_boilerplate/Components/Ticket"
	User "github.com/mahdidl/golang_boilerplate/Components/User"
	"github.com/mahdidl/golang_boilerplate/Components/UserRole"

	"net/http"
)

const (
	prefix               = "/api/v1"
	userPrefix           = "/user"
	ticketPrefix         = "/ticket"
	authenticationPrefix = "/auth"
	permissionPrefix     = "/permission"
	rolePrefix           = "/role"
	rolePermissionPrefix = "/role_permission"
	userRolePrefix       = "/user_role"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)
	routerTicket := app.Group(prefix + ticketPrefix)
	routerUser := app.Group(prefix + userPrefix)
	authUser := app.Group(prefix + authenticationPrefix)
	authPermission := app.Group(prefix + permissionPrefix)
	authRole := app.Group(prefix + rolePrefix)
	RolePermissionRouter := app.Group(prefix + rolePermissionPrefix)
	UserRoleRouter := app.Group(prefix + userRolePrefix)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	newUserRepository := User.NewUserRepository()
	newUserService := User.NewUserService(newUserRepository)
	newUserController := User.NewUserController(newUserService)

	newAuthUserRepository := AuthUser.NewAuthUserRepository()
	newAuthUserService := AuthUser.NewAuthUserService(newAuthUserRepository)
	newAuthUserController := AuthUser.NewAuthUserController(newAuthUserService)

	newTicketRepository := Ticket.NewTicketRepository()
	newTicketService := Ticket.NewTicketService(newUserService, newTicketRepository)
	newTicketController := Ticket.NewTicketController(newTicketService)

	newAuthService := Auth.NewAuthService()
	newAuthController := Auth.NewAuthController(newAuthService)

	newPermissionRepository := Permission.NewPermissionRepository()
	newPermissionService := Permission.NewPermissionService(newPermissionRepository)
	newPermissionController := Permission.NewPermissionController(newPermissionService)

	newRoleRepository := Role.NewRoleRepository()
	newRoleService := Role.NewRoleService(newRoleRepository)
	newRoleController := Role.NewRoleController(newRoleService)

	newRolePermissionRepository := RolePermission.NewRolePermissionRepository()
	newRolePermissionService := RolePermission.NewRolePermissionService(newRolePermissionRepository)
	newRolePermissionController := RolePermission.NewRolePermissionController(newRolePermissionService)

	newUserRoleRepository := UserRole.NewUserRoleRepository()
	newUserRoleService := UserRole.NewUserRoleService(newUserRoleRepository)
	newUserRoleController := UserRole.NewUserRoleController(newUserRoleService)

	// implement middleware to routes
	authTicketRoutes := routerTicket.Group("/").Use(Middleware.AuthMiddleware())
	authUserRoutes := routerUser.Group("/").Use(Middleware.AuthMiddleware())
	authPermissionRoutes := authPermission.Group("/").Use(Middleware.AuthMiddleware())
	authRoleRoutes := authRole.Group("/").Use(Middleware.AuthMiddleware())
	RolePermissionRoutes := RolePermissionRouter.Group("/").Use(Middleware.AuthMiddleware())
	UserRoleRoutes := UserRoleRouter.Group("/").Use(Middleware.AuthMiddleware())

	// user endpoints without auth
	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newUserController.LoginUser)
	routerUser.POST("/logout", newAuthUserController.Logout)

	// user endpoints with auth
	authUserRoutes.GET("/get_all", newUserController.GetAllUsers)
	authUserRoutes.GET("/", newUserController.GetUserById)
	authUserRoutes.PATCH("/", newUserController.UpdateUser)
	authUserRoutes.PATCH("/change_password", newUserController.ChangePassword)
	authUserRoutes.PATCH("/change_status", newUserController.ChangeActiveStatus)

	// authUser endpoints
	authUser.POST("/newToken", newAuthController.AccessToken)

	// ticket endpoints with auth
	authTicketRoutes.POST("/create", newTicketController.CreateTicket)

	// permission endpoints with auth
	// todo: remove create permission
	authPermissionRoutes.POST("/create", newPermissionController.CreatePermission)
	authPermissionRoutes.GET("/", newPermissionController.GetPermissions)

	// role endpoint with auth
	// create role
	authRoleRoutes.POST("/create", newRoleController.CreateRole)
	// get all roles
	authRoleRoutes.GET("/get_all", newRoleController.GetAllRoles)
	// get role with id
	authRoleRoutes.GET("/", newRoleController.GetRole)
	// update role with id
	authRoleRoutes.PATCH("/", newRoleController.UpdateRole)
	// delete role with id
	authRoleRoutes.DELETE("/", newRoleController.DeleteRole)

	// attach permission to role with roleId and permissionId
	RolePermissionRoutes.PATCH("/attach", newRolePermissionController.Attach)

	// detach permission from role with roleId and permissionId
	RolePermissionRoutes.PATCH("/detach", newRolePermissionController.Detach)

	// attach role to user with roleId and userId
	UserRoleRoutes.PATCH("/attach", newUserRoleController.Attach)

	// detach role from user with userId
	UserRoleRoutes.PATCH("/detach", newUserRoleController.Detach)

}
