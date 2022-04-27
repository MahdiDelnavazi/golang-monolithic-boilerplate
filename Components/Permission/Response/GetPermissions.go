package Entity

import "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"

type GetPermissions struct {
	Permissions []Entity.Permission `json:"Permissions"`
}
