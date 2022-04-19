package User

type AccessTokenRequest struct {
	AccessToken string `json:"refreshToken"  validate:"required"`
}
