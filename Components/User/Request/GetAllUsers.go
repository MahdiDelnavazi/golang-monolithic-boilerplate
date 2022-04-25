package Request

type GetAllUsers struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
}
