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

	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware

	"net/http"
)

const (
	prefix                = "/api/v1"
	usersPostfix          = "/user"
	ticketPostfix         = "/ticket"
	authenticationPostfix = "/authentication"
	permissionPostfix     = "/permission"
	rolePostfix           = "/role"
	rolePermissionPostfix = "/role-permission"
	userRolePostfix       = "/user-role"
)

func Routes(app *gin.Engine) {
	router := app.Group(prefix)
	//routerTicket := app.Group(prefix + ticketPostfix)
	//authUser := app.Group(prefix + authenticationPostfix)
	//authPermission := app.Group(prefix + permissionPostfix)
	//authRole := app.Group(prefix + rolePostfix)
	//RolePermissionRouter := app.Group(prefix + rolePermissionPostfix)
	//UserRoleRouter := app.Group(prefix + userRolePostfix)

	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userRepository := User.NewUserRepository()
	userService := User.NewUserService(userRepository)
	userController := User.NewUserController(userService)
	userRouter := router.Group(usersPostfix).Use(Middleware.AuthMiddleware())
	{
		// Get Requests
		userRouter.GET("", userController.GetAllUsers)
		userRouter.GET("/:userId", userController.GetUserById)

		// Put Requests
		userRouter.PUT("/:userId", userController.UpdateUser)
		userRouter.PUT("/:userId/password", userController.ChangePassword)

		// Patch Requests
		userRouter.PATCH("/:userId", userController.ChangeActiveStatus)
	}
	userRouter = router.Group(usersPostfix)
	{
		userRouter.POST("/create", userController.CreateUser)
	}

	authUserRepository := AuthUser.NewAuthUserRepository()
	authUserService := AuthUser.NewAuthUserService(authUserRepository)
	authUserController := AuthUser.NewAuthUserController(authUserService)

	authRepository := Auth.NewAuthRepository()
	authService := Auth.NewAuthService(authRepository)
	authController := Auth.NewAuthController(authService)

	authUserRouter := router.Group(authenticationPostfix).Use(Middleware.AuthMiddleware())
	{
		// Post Request
		authUserRouter.POST("/newToken", authController.AccessToken)

		// Delete Requests
		authUserRouter.DELETE("/logout", authUserController.Logout)
	}
	authUserRouter = router.Group(authenticationPostfix).Use()
	authUserRouter.POST("/login", authController.LoginUser)

	ticketRepository := Ticket.NewTicketRepository()
	ticketService := Ticket.NewTicketService(userService, ticketRepository)
	ticketController := Ticket.NewTicketController(ticketService)
	tickerRouter := router.Group(ticketPostfix).Use(Middleware.AuthMiddleware())
	{
		// Post Requests
		tickerRouter.POST("/create/:userId", ticketController.CreateTicket)
	}

	permissionRepository := Permission.NewPermissionRepository()
	permissionService := Permission.NewPermissionService(permissionRepository)
	permissionController := Permission.NewPermissionController(permissionService)
	permissionRouter := router.Group(permissionPostfix).Use(Middleware.AuthMiddleware())
	{
		// todo: remove create permission
		// Post Requests
		permissionRouter.POST("/", permissionController.CreatePermission)

		// Get Requests
		permissionRouter.GET("/", permissionController.GetPermissions)
	}

	// todo implement REST
	roleRepository := Role.NewRoleRepository()
	roleService := Role.NewRoleService(roleRepository)
	roleController := Role.NewRoleController(roleService)
	roleRouter := router.Group(rolePostfix).Use(Middleware.AuthMiddleware())
	{
		// Post Requests
		roleRouter.POST("/", roleController.CreateRole)

		// Get Requests
		roleRouter.GET("/", roleController.GetAllRoles)
		roleRouter.GET("/:roleId", roleController.GetRole)

		roleRouter.PATCH("/:roleId", roleController.UpdateRole)

		roleRouter.DELETE("/:roleId", roleController.DeleteRole)
	}

	rolePermissionRepository := RolePermission.NewRolePermissionRepository()
	rolePermissionService := RolePermission.NewRolePermissionService(rolePermissionRepository)
	rolePermissionController := RolePermission.NewRolePermissionController(rolePermissionService)
	rolePermissionRouter := router.Group(rolePermissionPostfix).Use(Middleware.AuthMiddleware())
	{
		// Patch Requests
		// attach permission to role with roleId and permissionId
		rolePermissionRouter.PATCH("/attach/:roleId/:permissionId", rolePermissionController.Attach)

		// detach permission from role with roleId and permissionId
		rolePermissionRouter.PATCH("/detach/:roleId/:permissionId", rolePermissionController.Detach)
	}

	userRoleRepository := UserRole.NewUserRoleRepository()
	userRoleService := UserRole.NewUserRoleService(userRoleRepository)
	userRoleController := UserRole.NewUserRoleController(userRoleService)
	userRoleRouter := router.Group(userRolePostfix).Use(Middleware.AuthMiddleware())
	{
		// Patch Requests
		// attach role to user with roleId and userId
		userRoleRouter.PATCH("/attach/:roleId/:userId", userRoleController.Attach)

		// detach role from user with userId
		userRoleRouter.PATCH("/detach/:userId", userRoleController.Detach)

	}

}
