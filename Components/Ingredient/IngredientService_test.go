package Ingredient

import (
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/Ingredient/Request"
	"github.com/mahdidl/golang_boilerplate/Test"
	"github.com/stretchr/testify/require"
	"testing"
)

var ingredientService *IngredientService

func init() {
	Test.OpenTestingDatabase()
	ingredientService = NewIngredientService(NewIngredientRepository())
}

func TestIngredientService_CreateIngredient(t *testing.T) {
	createIngredientRequest := Request.CreateIngredientRequest{Name: Helper.RandomString(5)}
	ingredient, err := ingredientService.CreateIngredient(createIngredientRequest.Name)
	require.NoError(t, err)
	require.NotNil(t, ingredient)
	require.NotEmpty(t, ingredient)
}

func TestIngredientService_GetAllIngredients(t *testing.T) {
	createIngredientRequest := Request.CreateIngredientRequest{Name: Helper.RandomString(5)}
	ingredient, err := ingredientService.CreateIngredient(createIngredientRequest.Name)
	require.NoError(t, err)
	require.NotNil(t, ingredient)
	require.NotEmpty(t, ingredient)

	getAllIngredientRequest := Request.GetAllIngredientRequest{Name: ""}
	ingredients, errGetAll := ingredientService.GetAllIngredients(getAllIngredientRequest.Name)
	require.NoError(t, errGetAll)
	require.NotNil(t, ingredients[0])
	require.NotEmpty(t, ingredients[0])
}
