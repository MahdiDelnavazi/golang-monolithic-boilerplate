package RolePermission

import (
	"errors"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	PermissionEntity "github.com/mahdidl/golang_boilerplate/Components/Permission/Entity"
	RoleEntity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
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

func (rolePermissionRepository RolePermissionRepository) Attach(roleId string, permissionId string) (RoleEntity.Role, error) {
	var role RoleEntity.Role
	var permission PermissionEntity.Permission

	PermissionId, err := primitive.ObjectIDFromHex(permissionId)
	if err = Config.PermissionCollection.FindOne(Config.DBContext, bson.M{"_id": PermissionId}).Decode(&permission); err != nil {
		return RoleEntity.Role{}, errors.New("permission not found")
	}

	RoleId, err := primitive.ObjectIDFromHex(roleId)

	update := bson.M{
		"$set": bson.M{
			"UpdatedAt": time.Now(),
		},
		"$addToSet": bson.M{
			"Permissions": PermissionId,
		},
	}

	resultErr := Config.RoleCollection.FindOneAndUpdate(Config.DBContext, bson.M{"_id": RoleId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)
	if resultErr != nil {
		return RoleEntity.Role{}, resultErr
	}

	return role, nil
}

func (rolePermissionRepository RolePermissionRepository) Detach(roleId string, permissionId string) (RoleEntity.Role, error) {
	var role RoleEntity.Role
	var permission PermissionEntity.Permission

	PermissionId, err := primitive.ObjectIDFromHex(permissionId)
	if err = Config.PermissionCollection.FindOne(Config.DBContext, bson.M{"_id": PermissionId}).Decode(&permission); err != nil {
		return RoleEntity.Role{}, errors.New("permission not found")
	}

	RoleId, err := primitive.ObjectIDFromHex(roleId)

	update := bson.M{
		"$pull": bson.M{
			"Permissions": PermissionId,
		},
	}

	resultErr := Config.RoleCollection.FindOneAndUpdate(Config.DBContext, bson.M{"_id": RoleId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)
	if resultErr != nil {
		return RoleEntity.Role{}, resultErr
	}

	return role, nil
}
