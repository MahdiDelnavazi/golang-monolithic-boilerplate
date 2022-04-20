package Response

import "golang_monolithic_bilerplate/Components/User/Entity"

type ResponseAllUsers struct {
	Users []Entity.User `json:"users" `
}
