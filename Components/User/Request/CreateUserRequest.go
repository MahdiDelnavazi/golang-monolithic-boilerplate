package Request

// swagger:parameters user createUser
type CreateUserRequest struct {
	// username of the user
	// in: string
	UserName string `json:"username" validate:"required,min=3"`
	// password of the user
	// in: string
	Password string `json:"password" validate:"required,min=8"`
}
