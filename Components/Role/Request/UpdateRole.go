package Request

type UpdateRole struct {
	Id   string `json:"Id" form:"Id" validate:"required,min=3"`
	Name string `json:"Name" bson:"Name"  validate:"required"`
}
