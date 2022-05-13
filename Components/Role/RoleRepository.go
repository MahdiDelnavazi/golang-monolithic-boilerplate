package Role

import (
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	Entity "github.com/mahdidl/golang_boilerplate/Components/Role/Entity"
	Request "github.com/mahdidl/golang_boilerplate/Components/Role/Request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type RoleRepository struct {
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

func (roleRepository *RoleRepository) Create(request Request.CreateRole) (Entity.Role, error) {
	role := Entity.Role{}
	slicePermissionId := make([]primitive.ObjectID, 0)

	result, err := Config.RoleCollection.InsertOne(Config.DBCtx, Entity.Role{Id: primitive.NewObjectID(), Name: request.Name, CreatedAt: time.Now(), PermissionsId: slicePermissionId})
	if err != nil {
		return Entity.Role{}, err
	}

	if err = Config.RoleCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&role); err != nil {
		return Entity.Role{}, err
	}

	return role, err
}

func (roleRepository *RoleRepository) Get(request Request.GetAllRole) ([]Entity.Role, error) {
	var roles = make([]Entity.Role, 0)

	roleCursor, queryError := Config.RoleCollection.Find(Config.DBCtx, bson.M{}, Helper.NewMongoPaginate(request.Limit, request.Page).GetPaginatedOpts())
	if queryError != nil {
		return nil, queryError
	}

	// decode permission and append to list
	for roleCursor.Next(Config.DBCtx) {
		var role Entity.Role
		if err := roleCursor.Decode(&role); err != nil {
			log.Println(err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (roleRepository RoleRepository) GetRoleById(Id string) (Entity.Role, error) {
	var role Entity.Role

	objectiveId, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		return Entity.Role{}, fmt.Errorf("id is not valid")
	}

	queryError := Config.RoleCollection.FindOne(Config.DBCtx, bson.M{"_id": objectiveId}).Decode(&role)
	if queryError != nil {
		return Entity.Role{}, fmt.Errorf("role not found")
	}

	return role, nil
}

func (roleRepository RoleRepository) Update(request Request.UpdateRole, roleId string) (Entity.Role, error) {
	var role Entity.Role

	id1, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		return Entity.Role{}, fmt.Errorf("id is not valid")
	}
	update := bson.D{
		{"$set", bson.D{{"Name", request.Name}, {"UpdatedAt", time.Now()}}},
	}

	result := Config.RoleCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": id1}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)

	if result != nil {
		return Entity.Role{}, fmt.Errorf("role not found")
	}
	return role, result
}

func (roleRepository RoleRepository) Delete(roleId string) (Entity.Role, error) {
	var role Entity.Role

	id1, err := primitive.ObjectIDFromHex(roleId)
	if err != nil {
		return Entity.Role{}, fmt.Errorf("id is not valid")
	}
	update := bson.D{
		{"$set", bson.D{{"DeletedAt", time.Now()}}},
	}

	result := Config.RoleCollection.FindOneAndUpdate(Config.DBCtx, bson.M{"_id": id1}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&role)

	if result != nil {
		return Entity.Role{}, fmt.Errorf("role not found")
	}
	return role, result
}
