package Request

type DetachRole struct {
	UserId string `json:"userId" form:"userId" validate:"required,min=3"`
}
