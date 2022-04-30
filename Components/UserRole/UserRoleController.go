package UserRole

import (
	"github.com/gin-gonic/gin"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"

	Request "github.com/mahdidl/golang_boilerplate/Components/UserRole/Request"
	Response "github.com/mahdidl/golang_boilerplate/Components/UserRole/Response"
	"net/http"
)

type UserRoleController struct {
	userRoleService *UserRoleService
}

func NewUserRoleController(userRoleService *UserRoleService) *UserRoleController {
	return &UserRoleController{userRoleService: userRoleService}
}

func (userRoleController *UserRoleController) Attach(context *gin.Context) {
	var request Request.AttachRole
	context.ShouldBindQuery(&request)

	user, responseErr := userRoleController.userRoleService.AttachRole(request)

	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: Response.AttachRole{ID: user.ID, CreatedAt: user.CreatedAt, RoleID: user.RoleID, UserName: user.UserName, IsActive: user.IsActive,
		UpdatedAt: user.UpdatedAt, DeletedAt: user.DeletedAt}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (userRoleController *UserRoleController) Detach(context *gin.Context) {
	var request Request.DetachRole
	context.ShouldBindQuery(&request)

	user, responseErr := userRoleController.userRoleService.DetachRole(request)
	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: Response.DetachRole{ID: user.ID, CreatedAt: user.CreatedAt, RoleID: user.RoleID, UserName: user.UserName, IsActive: user.IsActive,
		UpdatedAt: user.UpdatedAt, DeletedAt: user.DeletedAt}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
