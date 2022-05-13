package Permission

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	_ "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"net/http"
)

type PermissionController struct {
	PermissionService *PermissionService
}

func NewPermissionController(service *PermissionService) *PermissionController {
	return &PermissionController{PermissionService: service}
}

// @Summary      Create permission
// @Description  Create permission
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        createPermissionRequest  body      Request.CreatePermission  true  "Create permission request"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Permission}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "name should have more than 3 character"
// @Router       /permission [post]
// @Security ApiKeyAuth
//
// CreatePermission is a handler function which is creating new permission
func (permissionController *PermissionController) CreatePermission(context *gin.Context) {
	var createPermission Request.CreatePermission

	validationError := context.ShouldBindJSON(&createPermission)
	if validationError != nil {
		response := Response.GeneralResponse{
			Error: true, Message: validationError.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	validateError := Validator.ValidationCheck(createPermission)
	if validateError != nil {
		response := Response.GeneralResponse{
			Error: true, Message: validateError.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	permisssionResponse, responseErr := permissionController.PermissionService.CreateNewPermission(createPermission)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: permisssionResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// @Summary      Get permission
// @Description  Get permission
// @Tags         Permission
// @Accept       json
// @Produce      json
// @Param        getAllPermissionsRequest  query      Request.GetAllPermissions  true  "get permissions"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Permission}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /permission [get]
// @Security ApiKeyAuth
//
// GetPermissions is a handler function which is return permissions with pagination
func (permissionController *PermissionController) GetPermissions(context *gin.Context) {
	var request Request.GetAllPermissions
	context.ShouldBindQuery(&request)

	permissionResponse, responseErr := permissionController.PermissionService.GetPermissions(request)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: permissionResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
