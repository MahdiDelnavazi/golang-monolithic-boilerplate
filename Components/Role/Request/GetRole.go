package Request

type GetRole struct {
	Id string `json:"roleId" form:"roleId" validate:"required,min=3"`
}
