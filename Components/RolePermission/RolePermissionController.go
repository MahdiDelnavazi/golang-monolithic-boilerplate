package RolePermission

import (
	"github.com/gin-gonic/gin"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"
	_ "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"net/http"
)

type RolePermissionController struct {
	RolePermissionService *RolePermissionService
}

func NewRolePermissionController(rolePermissionService *RolePermissionService) *RolePermissionController {
	return &RolePermissionController{RolePermissionService: rolePermissionService}
}

// @Summary      Attach permission to role
// @Description  Attach permission to role with roleId and permissionId
// @Tags         Role-Permission
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "roleId"
// @Param        permissionId  path      string  true  "permissionId"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Role}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /role-permission/attach/{roleId}{permissionId} [patch]
// @Security ApiKeyAuth
//
// Attach permission to role
func (rolePermissionController *RolePermissionController) Attach(context *gin.Context) {
	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "roleId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	permissionId := context.Param("permissionId")
	validationErr = primitive.IsValidObjectID(permissionId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "permissionId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	role, responseErr := rolePermissionController.RolePermissionService.Attach(roleId, permissionId)
	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: role}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// @Summary      Detach permission from role
// @Description  Detach permission from role with roleId and permissionId
// @Tags         Role-Permission
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "roleId"
// @Param        permissionId  path      string  true  "permissionId"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Role}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /role-permission/detach/{roleId}{permissionId} [patch]
// @Security ApiKeyAuth
//
// Detach permission from role
func (rolePermissionController *RolePermissionController) Detach(context *gin.Context) {
	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "roleId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	permissionId := context.Param("permissionId")
	validationErr = primitive.IsValidObjectID(permissionId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "permissionId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	role, responseErr := rolePermissionController.RolePermissionService.Detach(roleId, permissionId)
	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: role}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
