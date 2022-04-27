package Request

type DeleteRole struct {
	Id string `json:"Id" form:"Id" validate:"required,min=3"`
}
