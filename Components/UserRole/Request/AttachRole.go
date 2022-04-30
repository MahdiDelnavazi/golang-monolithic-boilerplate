package Request

type AttachRole struct {
	RoleId string `json:"roleId" form:"roleId" validate:"required,min=3"`
	UserId string `json:"userId" form:"userId" validate:"required,min=3"`
}
