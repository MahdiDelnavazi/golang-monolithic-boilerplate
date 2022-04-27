package Response

import "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"

type GetRole struct {
	Roles Entity.Role `json:"Role"`
}
