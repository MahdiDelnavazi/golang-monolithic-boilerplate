package Response

type LoginUserResponse struct {
	Id           string `json:"id" `
	UserName     string `json:"username" `
	AccessToken  string `json:"accessToken" `
	RefreshToken string `json:"refreshToken" `
}
