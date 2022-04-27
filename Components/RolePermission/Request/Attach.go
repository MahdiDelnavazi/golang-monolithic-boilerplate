package Request

type AttachPermission struct {
	RoleId       string `json:"roleId" form:"roleId" validate:"required,min=3"`
	PermissionId string `json:"permissionId" form:"permissionId" validate:"required,min=3"`
}
