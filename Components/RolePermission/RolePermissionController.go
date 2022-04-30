package RolePermission

import (
	"github.com/gin-gonic/gin"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Components/RolePermission/Request"
	Response "github.com/mahdidl/golang_boilerplate/Components/RolePermission/Response"

	"net/http"
)

type RolePermissionController struct {
	RolePermissionService *RolePermissionService
}

func NewRolePermissionController(rolePermissionService *RolePermissionService) *RolePermissionController {
	return &RolePermissionController{RolePermissionService: rolePermissionService}
}

func (rolePermissionController *RolePermissionController) Attach(context *gin.Context) {
	var request Request.AttachPermission
	context.ShouldBindQuery(&request)

	role, responseErr := rolePermissionController.RolePermissionService.Attach(request)
	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: Response.Attach{Id: role.Id, Name: role.Name, PermissionsId: role.PermissionsId, CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt, DeletedAt: role.DeletedAt}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (rolePermissionController *RolePermissionController) Detach(context *gin.Context) {
	var request Request.DetachPermission
	context.ShouldBindQuery(&request)

	role, responseErr := rolePermissionController.RolePermissionService.Detach(request)
	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: Response.Attach{Id: role.Id, Name: role.Name, PermissionsId: role.PermissionsId, CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt, DeletedAt: role.DeletedAt}}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
