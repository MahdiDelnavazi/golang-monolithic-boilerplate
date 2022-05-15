package Request

type GetAllIngredientRequest struct {
	Name string `form:"name" validate:"max=100"`
}
