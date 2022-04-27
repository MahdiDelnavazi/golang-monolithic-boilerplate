package Permission

import (
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	Entity "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	Request "github.com/mahdidl/golang_boilerplate/Components/Permission/Request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type PermissionRepository struct {
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

func (permissionRepository *PermissionRepository) CreateNewPermission(request Request.CreatePermission) (Entity.Permission, error) {
	permission := Entity.Permission{}

	result, err := Config.PermissionCollection.InsertOne(Config.DBCtx, Entity.Permission{Id: primitive.NewObjectID(), Name: request.Name})
	if err != nil {
		return Entity.Permission{}, err
	}

	if err = Config.PermissionCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&permission); err != nil {
		return Entity.Permission{}, err
	}

	return permission, err
}

func (permissionRepository *PermissionRepository) GetPermissions(request Request.GetAllPermissions) ([]Entity.Permission, error) {
	var permissions = make([]Entity.Permission, 0)

	permissionCursor, queryError := Config.PermissionCollection.Find(Config.DBCtx, bson.M{}, Helper.NewMongoPaginate(request.Limit, request.Page).GetPaginatedOpts())
	if queryError != nil {
		return nil, queryError
	}

	// decode permission and append to list
	for permissionCursor.Next(Config.DBCtx) {
		var permission Entity.Permission
		if err := permissionCursor.Decode(&permission); err != nil {
			log.Println(err)
		}
		permissions = append(permissions, permission)
	}

	return permissions, nil
}
