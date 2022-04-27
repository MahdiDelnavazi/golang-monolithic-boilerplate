package Response

import "github.com/mahdidl/golang_boilerplate/Components/User/Entity"

type ResponseAllUsers struct {
	Users []Entity.User `json:"users" `
}
