package Request

type UpdateRole struct {
	Name string `json:"Name" bson:"Name"  validate:"required"`
}
