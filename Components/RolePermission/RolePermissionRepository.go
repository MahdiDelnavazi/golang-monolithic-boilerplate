package RolePermission

import (
	"errors"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	PermissionEntity "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	RoleEntity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/RolePermission/Request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RolePermissionRepository struct {
}

func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (rolePermissionRepository RolePermissionRepository) Attach(request Request.AttachPermission) (RoleEntity.Role, error) {
	var role RoleEntity.Role
	var permission PermissionEntity.Permission

	PermissionId, err := primitive.ObjectIDFromHex(request.PermissionId)
	if err = Config.PermissionCollection.FindOne(Config.DBCtx, bson.M{"_id": PermissionId}).Decode(&permission); err != nil {
		return RoleEntity.Role{}, errors.New("permission not found")
	}

	RoleId, err := primitive.ObjectIDFromHex(request.RoleId)

	update := bson.M{
		"$set": bson.M{
			"UpdatedAt": time.Now(),
		},
		"$addToSet": bson.M{
			"Permissions": PermissionId,
		},
	}

	resultErr := Config.RoleCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": RoleId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)
	if resultErr != nil {
		return RoleEntity.Role{}, resultErr
	}

	return role, nil
}

func (rolePermissionRepository RolePermissionRepository) Detach(request Request.DetachPermission) (RoleEntity.Role, error) {
	var role RoleEntity.Role
	var permission PermissionEntity.Permission

	PermissionId, err := primitive.ObjectIDFromHex(request.PermissionId)
	if err = Config.PermissionCollection.FindOne(Config.DBCtx, bson.M{"_id": PermissionId}).Decode(&permission); err != nil {
		return RoleEntity.Role{}, errors.New("permission not found")
	}

	RoleId, err := primitive.ObjectIDFromHex(request.RoleId)

	update := bson.M{
		"$pull": bson.M{
			"Permissions": PermissionId,
		},
	}

	resultErr := Config.RoleCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": RoleId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)
	if resultErr != nil {
		return RoleEntity.Role{}, resultErr
	}

	return role, nil
}
