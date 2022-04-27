package Request

type GetAllRole struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}
