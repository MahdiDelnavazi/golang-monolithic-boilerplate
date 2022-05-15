package Ingredient

import (
	"errors"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Components/Ingredient/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type IngredientRepository struct {
}

func NewIngredientRepository() *IngredientRepository {
	return &IngredientRepository{}
}

// CreateIngredient is creating a new Ingredient and return it
func (ingredientRepository *IngredientRepository) CreateIngredient(name string) (Entity.Ingredient, error) {
	var ingredient Entity.Ingredient

	// check if ingredient is not exist
	Config.IngredientCollection.FindOne(Config.DBCtx, bson.M{"Name": name}).Decode(&ingredient)
	if ingredient.Name != "" {
		return Entity.Ingredient{}, errors.New("ingredient is exist")
	}

	result, err := Config.IngredientCollection.InsertOne(Config.DBCtx, Entity.Ingredient{Id: primitive.NewObjectID(), Name: name, CreatedAt: time.Now()})
	if err != nil {
		return Entity.Ingredient{}, err
	}

	err = Config.IngredientCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&ingredient)
	if err != nil {
		return Entity.Ingredient{}, err
	}

	return ingredient, err
}

// getAllIngredients returns all Ingredients with name search
func (ingredientRepository *IngredientRepository) getAllIngredients(filter string) ([]Entity.Ingredient, error) {
	var ingredientList = make([]Entity.Ingredient, 0)

	// if we have filter option we filter list with it , else we return whole list of ingredients
	var filterQuery = bson.M{}
	if filter != "" {
		filterQuery = bson.M{"Name": filter}
	}

	userCursor, queryError := Config.IngredientCollection.Find(Config.DBCtx, filterQuery)
	if queryError != nil {
		return nil, queryError
	}

	// decode users and append to list
	for userCursor.Next(Config.DBCtx) {
		var ingredient Entity.Ingredient
		if err := userCursor.Decode(&ingredient); err != nil {
			log.Println(err)
		}
		ingredientList = append(ingredientList, ingredient)
	}

	return ingredientList, nil
}
