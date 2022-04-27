package Response

type UpdateRole struct {
	Id   string `json:"Id" validate:"required,min=3"`
	Name string `json:"Name" bson:"Name"  validate:"required"`
}
