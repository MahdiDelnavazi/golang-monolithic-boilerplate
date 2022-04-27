package Request

type CreateRole struct {
	Name string `json:"Name" validate:"required,min=3"`
}
