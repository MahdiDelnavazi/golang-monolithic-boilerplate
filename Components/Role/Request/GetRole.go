package Request

type GetRole struct {
	Id string `json:"Id" form:"Id" validate:"required,min=3"`
}
