package Request

type ChangePasswordRequest struct {
	CurrentPassword string `bson:"Password" validate:"required,min=8"  json:"currentPassword"`
	NewPassword     string `bson:"Password" validate:"required,min=8" json:"newPassword"`
}
