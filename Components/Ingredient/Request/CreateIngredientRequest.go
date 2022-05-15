package Request

type CreateIngredientRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}
