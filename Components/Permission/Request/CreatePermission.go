package Request

type CreatePermission struct {
	Name string `json:"Name" validate:"required,min=3"`
}
