package User

type LogoutRequest struct {
	Token string `json:"token" validate:"required,min=3"`
}
