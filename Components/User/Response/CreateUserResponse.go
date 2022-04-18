package Response

type CreateUserResponse struct {
	UserName string `json:"username" binding:"required"`
}
