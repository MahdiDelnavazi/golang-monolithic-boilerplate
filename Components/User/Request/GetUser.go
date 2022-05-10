package Request

type GetUser struct {
	ID string `json:"userId" form:"userId"`
}
