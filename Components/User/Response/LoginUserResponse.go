package Response

type LoginUserResponse struct {
	UserName     string `json:"username" `
	ID           string `json:"id" `
	AccessToken  string `json:"accessToken" `
	RefreshToken string `json:"refreshToken" `
}
