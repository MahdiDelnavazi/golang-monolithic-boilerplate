package UserRole

import (
	"errors"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	RoleEntity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/UserRole/Entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRoleRepository struct {
}

func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (rolePermissionRepository UserRoleRepository) Attach(userId string, roleId string) (Entity.UserRole, error) {
	var role RoleEntity.Role
	var user Entity.UserRole

	objectUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid user id")
	}
	objectRoleId, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid role id")
	}

	if err = Config.RoleCollection.FindOne(Config.DBCtx, bson.M{"_id": objectRoleId}).Decode(&role); err != nil {
		return Entity.UserRole{}, errors.New("role not found")
	}

	update := bson.D{
		{"$set", bson.D{{"RoleId", objectRoleId}, {"UpdatedAt", time.Now()}}},
	}

	resultErr := Config.UserCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": objectUserId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if resultErr != nil {
		return Entity.UserRole{}, resultErr
	}

	return user, nil
}

func (rolePermissionRepository UserRoleRepository) Detach(userId string) (Entity.UserRole, error) {
	var user Entity.UserRole

	objectUserId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid user id")
	}
	update := bson.D{
		{"$set", bson.D{{"RoleId", nil}, {"UpdatedAt", time.Now()}}},
	}

	resultErr := Config.UserCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": objectUserId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if resultErr != nil {

		return Entity.UserRole{}, resultErr
	}

	return user, nil
}
