package UserRole

import (
	"github.com/gin-gonic/gin"
	General "github.com/mahdidl/golang_boilerplate/Common/Response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	Response "github.com/mahdidl/golang_boilerplate/Components/UserRole/Response"
	"net/http"
)

type UserRoleController struct {
	userRoleService *UserRoleService
}

func NewUserRoleController(userRoleService *UserRoleService) *UserRoleController {
	return &UserRoleController{userRoleService: userRoleService}
}

// @Summary      Atttach role from user
// @Description  Atttach role to user with roleId and userId
// @Tags         User-Role
// @Accept       json
// @Produce      json
// @Param        roleId  path      string  true  "roleId"
// @Param        userId  path      string  true  "userId"
// @Success      200                {object}  General.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  General.GeneralResponse{data=object} ""
// @Router       /user-role/attach/{roleId}{userId} [patch]
// @Security ApiKeyAuth
//
// Attach permission from role
func (userRoleController *UserRoleController) Attach(context *gin.Context) {
	roleId := context.Param("roleId")

	validationErr := primitive.IsValidObjectID(roleId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "roleId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	userId := context.Param("userId")
	validationErr = primitive.IsValidObjectID(userId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "userId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	user, responseErr := userRoleController.userRoleService.AttachRole(userId, roleId)

	if responseErr != nil {
		response := General.GeneralResponse{
			Error: true, Message: responseErr.Error(),
		}
		context.JSON(http.StatusInternalServerError, gin.H{"response": response})
		return
	}

	response := General.GeneralResponse{Error: false, Message: "successful", Data: user}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// @Summary      Dettach role from user
// @Description  Dettach role from user with roleId and userId
// @Tags         User-Role
// @Accept       json
// @Produce      json
// @Param        userId  path      string  true  "userId"
// @Success      200                {object}  General.GeneralResponse{data=Entity.User}
// @Failure      400                {object}  General.GeneralResponse{data=object} ""
// @Router       /user-role/detach/{userId} [patch]
// @Security ApiKeyAuth
//
// Detach permission from role
func (userRoleController *UserRoleController) Detach(context *gin.Context) {
	userId := context.Param("userId")
	validationErr := primitive.IsValidObjectID(userId)

	if !validationErr {
		response := General.GeneralResponse{Error: true, Message: "userId is not valid"}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	user, responseErr := userRoleController.userRoleService.DetachRole(userId)
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
