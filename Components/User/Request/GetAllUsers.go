package Request

type GetAllUsers struct {
	Limit int `json:"limit" `
	Page  int `json:"page"`
}
