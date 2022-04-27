package Request

type GetAllPermissions struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}
