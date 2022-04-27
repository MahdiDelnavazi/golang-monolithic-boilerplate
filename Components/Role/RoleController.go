package Role

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	Request "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"net/http"
)

type RoleController struct {
	RoleService *RoleService
}

func NewRoleController(service *RoleService) *RoleController {
	return &RoleController{RoleService: service}
}

func (roleController RoleController) CreateRole(context *gin.Context) {
	var createPermission Request.CreateRole

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

	permisssionResponse, responseErr := roleController.RoleService.Create(createPermission)
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

func (roleController *RoleController) GetAllRoles(context *gin.Context) {
	var request Request.GetAllRole
	context.ShouldBindQuery(&request)

	permissionResponse, responseErr := roleController.RoleService.GetAll(request)
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

func (roleController *RoleController) GetRole(context *gin.Context) {
	var request Request.GetRole
	context.ShouldBindQuery(&request)

	role, responseErr := roleController.RoleService.GetRoleById(request.Id)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: role}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (roleController *RoleController) UpdateRole(context *gin.Context) {
	var request Request.UpdateRole
	context.ShouldBindQuery(&request)
	context.ShouldBindJSON(&request)

	role, responseErr := roleController.RoleService.Update(request)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: role}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

func (roleController *RoleController) DeleteRole(context *gin.Context) {
	var request Request.DeleteRole
	context.ShouldBindQuery(&request)

	_, responseErr := roleController.RoleService.Delete(request)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: nil}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
