package UserRole

import (
	"errors"
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	RoleEntity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/UserRole/Entity"
	Request "github.com/mahdidl/golang_boilerplate/Components/UserRole/Request"
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

func (rolePermissionRepository UserRoleRepository) Attach(request Request.AttachRole) (Entity.UserRole, error) {
	var role RoleEntity.Role
	var user Entity.UserRole

	fmt.Println("salaam ", request)
	UserId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid user id")
	}
	RoleId, err := primitive.ObjectIDFromHex(request.RoleId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid role id")
	}

	if err = Config.RoleCollection.FindOne(Config.DBCtx, bson.M{"_id": RoleId}).Decode(&role); err != nil {
		return Entity.UserRole{}, errors.New("role not found")
	}

	update := bson.D{
		{"$set", bson.D{{"RoleId", RoleId}, {"UpdatedAt", time.Now()}}},
	}

	resultErr := Config.UserCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": UserId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if resultErr != nil {
		return Entity.UserRole{}, resultErr
	}

	return user, nil
}

func (rolePermissionRepository UserRoleRepository) Detach(request Request.DetachRole) (Entity.UserRole, error) {
	var user Entity.UserRole

	fmt.Println("this is bug ", request)
	userId, err := primitive.ObjectIDFromHex(request.UserId)
	if err != nil {
		return Entity.UserRole{}, errors.New("invalid user id")
	}
	update := bson.D{
		{"$set", bson.D{{"RoleId", nil}, {"UpdatedAt", time.Now()}}},
	}

	resultErr := Config.UserCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": userId}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if resultErr != nil {
		fmt.Println("this is result error  ", resultErr)
		return Entity.UserRole{}, resultErr
	}

	return user, nil
}
