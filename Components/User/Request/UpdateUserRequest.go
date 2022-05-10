package Request

type UpdateUserRequest struct {
	UserName string `bson:"UserName" json:"userName"`
}
