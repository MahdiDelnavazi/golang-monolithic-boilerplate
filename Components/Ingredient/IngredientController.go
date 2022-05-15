package Ingredient

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	"github.com/mahdidl/golang_boilerplate/Components/Ingredient/Request"
	"log"
	"net/http"
)

type IngredientController struct {
	ingredientService *IngredientService
}

func NewIngredientController(ingredientService *IngredientService) *IngredientController {
	return &IngredientController{ingredientService: ingredientService}
}

// @Summary      Create Ingredient
// @Description  Create Ingredient
// @Tags         Ingredient
// @Accept       json
// @Produce      json
// @Param        CreateIngredientRequest  body      Request.CreateIngredientRequest  true  "Create Ingreedient request"
// @Success      200                {object}  Response.GeneralResponse{data=Entity.Ingredient}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /ingredient [post]
// @Security ApiKeyAuth
//
// CreateUser is a handler function which is creating user
func (ingredientController *IngredientController) CreateIngredient(context *gin.Context) {
	var createIngredientRequest Request.CreateIngredientRequest

	context.ShouldBindJSON(&createIngredientRequest)

	validationError := Validator.ValidationCheck(createIngredientRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	ingredientResponse, responseError := ingredientController.ingredientService.CreateIngredient(createIngredientRequest.Name)

	if responseError != nil {
		response := Response.GeneralResponse{Error: true, Message: responseError.Error(), Data: nil}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "ingredient have been created", Data: ingredientResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}

// @Summary      Get All Ingredients
// @Description  Get All Ingredients
// @Tags         Ingredient
// @Accept       json
// @Produce      json
// @Param        name  query      string  false  "filter ingredients with name"
// @Success      200                {object}  Response.GeneralResponse{data=[]Entity.Ingredient}
// @Failure      400                {object}  Response.GeneralResponse{data=object} ""
// @Router       /ingredient [get]
// @Security ApiKeyAuth
//
// GetAllIngredient is a handler function which is return all Ingredients with filter in name (optional)
func (ingredientController *IngredientController) GetAllIngredient(context *gin.Context) {
	var getAllIngredientRequest Request.GetAllIngredientRequest
	context.ShouldBindQuery(&getAllIngredientRequest)

	validationError := Validator.ValidationCheck(getAllIngredientRequest)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	ingredientResponse, responseError := ingredientController.ingredientService.GetAllIngredients(getAllIngredientRequest.Name)

	if responseError != nil {
		response := Response.GeneralResponse{Error: true, Message: responseError.Error(), Data: nil}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
		return
	}

	// all ok
	// create general response
	response := Response.GeneralResponse{Error: false, Message: "ingredient list", Data: ingredientResponse}
	context.JSON(http.StatusOK, gin.H{"response": response})
}
