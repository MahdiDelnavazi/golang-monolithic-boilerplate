package Response

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"  validate:"required"`
}
