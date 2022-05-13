package Role

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	_ "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	Request "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	RoleResponse "github.com/mahdidl/golang_boilerplate/Components/Role/Response"
	_ "github.com/mahdidl/golang_boilerplate/Components/Role/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RoleController struct {
	RoleService *RoleService
}

func NewRoleController(service *RoleService) *RoleController {
	return &RoleController{RoleService: service}
}

// @Summary      Create role
// @Description  Create role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        CreateRoleRequest  body      Request.CreateRole  true  "Create role request"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Role}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "create role"
// @Router       /role/ [post]
// @Security ApiKeyAuth
//
// CreateRole is a handler function which is creating role
func (roleController RoleController) CreateRole(context *gin.Context) {
	var createRole Request.CreateRole

	validationError := context.ShouldBindJSON(&createRole)
	if validationError != nil {
		response := Response.GeneralResponse{
			Error: true, Message: validationError.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	validateError := Validator.ValidationCheck(createRole)
	if validateError != nil {
		response := Response.GeneralResponse{
			Error: true, Message: validateError.Error(),
		}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	permisssionResponse, responseErr := roleController.RoleService.Create(createRole)
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

// @Summary      Create role
// @Description  Create role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        GetAllRoleRequest  query      Request.GetAllRole  true  "get all roles with pagination"
// @Success      200                {object}  Response.GeneralResponse{data=RoleResponse.GetAllRoles}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "create role"
// @Router       /role/ [get]
// @Security ApiKeyAuth
//
// GetAllRoles is a handler function which is return all roles with pagination
func (roleController *RoleController) GetAllRoles(context *gin.Context) {
	var request Request.GetAllRole
	context.ShouldBindQuery(&request)
	roleResponse, responseErr := roleController.RoleService.GetAll(request)
	if responseErr != nil {
		response := Response.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	var role RoleResponse.GetAllRoles
	role = roleResponse

	response := Response.GeneralResponse{Error: false, Message: "successful", Data: role.Roles}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// @Summary      Get role
// @Description  Get role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "get role with id"
// @Success      200                {object}  Response.GeneralResponse{data=RoleResponse.GetRole}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "get role"
// @Router       /role/{roleId} [get]
// @Security ApiKeyAuth
//
// GetRole is a handler function which is return role
func (roleController *RoleController) GetRole(context *gin.Context) {

	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)
	fmt.Println("this is role id :", roleId, validationErr)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	role, responseErr := roleController.RoleService.GetRoleById(roleId)
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

// @Summary      Get role
// @Description  Get role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "update role with id"
// @Param        UpdateRole  body      Request.UpdateRole  true  "update role model"
// @Success      200                {object}  Response.GeneralResponse{data=RoleResponse.GetRole}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "get role"
// @Router       /role/{roleId} [patch]
// @Security ApiKeyAuth
//
// UpdateRole for updating role name
func (roleController *RoleController) UpdateRole(context *gin.Context) {
	var request Request.UpdateRole
	context.ShouldBindJSON(&request)

	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)
	fmt.Println("this is role id :", roleId, validationErr)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	role, responseErr := roleController.RoleService.Update(request, roleId)
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

// @Summary      Delete role
// @Description  Delete role
// @Tags         Role
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "delete role with id"
// @Success      200                {object}  Response.GeneralResponse{data=RoleResponse.GetRole}
// @Failure      400                {object}  Response.GeneralResponse{data=object} "get role"
// @Router       /role/{roleId} [delete]
// @Security ApiKeyAuth
//
// UpdateRole for updating role name
func (roleController *RoleController) DeleteRole(context *gin.Context) {
	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)
	fmt.Println("this is role id :", roleId, validationErr)

	if !validationErr {
		response := Response.GeneralResponse{Error: true, Message: "id is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	_, responseErr := roleController.RoleService.Delete(roleId)
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
