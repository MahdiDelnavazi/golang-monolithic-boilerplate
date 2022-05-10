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
	"github.com/mahdidl/golang_boilerplate/docs"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
)

const (
	prefix               = "/api/v1"
	userPrefix           = "/user"
	usersPrefix          = "/users"
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
	routerUsers := app.Group(prefix + usersPrefix)
	authUser := app.Group(prefix + authenticationPrefix)
	authPermission := app.Group(prefix + permissionPrefix)
	authRole := app.Group(prefix + rolePrefix)
	RolePermissionRouter := app.Group(prefix + rolePermissionPrefix)
	UserRoleRouter := app.Group(prefix + userRolePrefix)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	newAuthRepository := Auth.NewAuthRepository()
	newAuthService := Auth.NewAuthService(newAuthRepository)
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
	authUsersRoutes := routerUsers.Group("/").Use(Middleware.AuthMiddleware())
	authPermissionRoutes := authPermission.Group("/").Use(Middleware.AuthMiddleware())
	authRoleRoutes := authRole.Group("/").Use(Middleware.AuthMiddleware())
	RolePermissionRoutes := RolePermissionRouter.Group("/").Use(Middleware.AuthMiddleware())
	UserRoleRoutes := UserRoleRouter.Group("/").Use(Middleware.AuthMiddleware())

	// user endpoints without auth
	routerUser.POST("/create", newUserController.CreateUser)
	routerUser.POST("/login", newAuthController.LoginUser)

	//done
	authUser.POST("/logout", newAuthUserController.Logout)

	//	done
	// user endpoints with auth
	authUsersRoutes.GET("/", newUserController.GetAllUsers)
	authUserRoutes.GET("/:userId", newUserController.GetUserById)
	authUserRoutes.PATCH("/:userId", newUserController.UpdateUser)
	authUserRoutes.PUT("/change-password/:userId", newUserController.ChangePassword)
	authUserRoutes.PATCH("/change-status/:userId", newUserController.ChangeActiveStatus)

	// done
	// authUser endpoints
	authUser.POST("/newToken", newAuthController.AccessToken)

	//	done
	// ticket endpoints with auth
	authTicketRoutes.POST("/create/:userId", newTicketController.CreateTicket)

	// permission endpoints with auth
	// todo: remove create permission
	authPermissionRoutes.POST("/create", newPermissionController.CreatePermission)
	authPermissionRoutes.GET("/", newPermissionController.GetPermissions)

	// role endpoint with auth
	// create role
	authRoleRoutes.POST("/create", newRoleController.CreateRole)
	// get all roles
	authRoleRoutes.GET("/get-all", newRoleController.GetAllRoles)
	// get role with id
	authRoleRoutes.GET("/:roleId", newRoleController.GetRole)
	// update role with id
	authRoleRoutes.PATCH("/:roleId", newRoleController.UpdateRole)
	// delete role with id
	authRoleRoutes.DELETE("/:roleId", newRoleController.DeleteRole)

	// attach permission to role with roleId and permissionId
	RolePermissionRoutes.PATCH("/attach", newRolePermissionController.Attach)

	// detach permission from role with roleId and permissionId
	RolePermissionRoutes.PATCH("/detach", newRolePermissionController.Detach)

	// attach role to user with roleId and userId
	UserRoleRoutes.PATCH("/attach", newUserRoleController.Attach)

	// detach role from user with userId
	UserRoleRoutes.PATCH("/detach", newUserRoleController.Detach)

}
