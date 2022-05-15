package Ingredient

import "github.com/mahdidl/golang_boilerplate/Components/Ingredient/Entity"

type IngredientService struct {
	ingredientRepository *IngredientRepository
}

func NewIngredientService(ingredientRepository *IngredientRepository) *IngredientService {
	return &IngredientService{ingredientRepository: ingredientRepository}
}

func (ingredientService *IngredientService) CreateIngredient(name string) (Entity.Ingredient, error) {
	ingredient, err := ingredientService.ingredientRepository.CreateIngredient(name)
	if err != nil {
		return Entity.Ingredient{}, err
	}

	return ingredient, nil
}

func (ingredientService *IngredientService) GetAllIngredients(filter string) ([]Entity.Ingredient, error) {
	ingredient, err := ingredientService.ingredientRepository.getAllIngredients(filter)
	if err != nil {
		return []Entity.Ingredient{}, err
	}

	return ingredient, nil
}
