package Request

type ChangePasswordRequest struct {
	Id              string `bson:"_id" json:"id" form:"id"`
	CurrentPassword string `bson:"Password" validate:"required,min=8"  json:"currentPassword"`
	NewPassword     string `bson:"Password" validate:"required,min=8" json:"newPassword"`
}
