package Controller

import (
	"fmt"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Components/User/Entity"
	"github.com/mahdidl/golang_boilerplate/Components/User/Request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser exec query for create new user in database
func (userRepository *UserRepository) CreateUser(creatUserRequest Request.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}

	result, err := Config.UserCollection.InsertOne(Config.DBContext, Entity.User{ID: primitive.NewObjectID(), IsActive: true,
		UserName: creatUserRequest.UserName, Password: creatUserRequest.Password, CreatedAt: time.Now()})
	if err != nil {
		return Entity.User{}, err
	}

	if err = Config.UserCollection.FindOne(Config.DBContext, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		return Entity.User{}, err
	}

	return user, err
}

// LoginUser for login users
func (userRepository *UserRepository) LoginUser(loginUserRequest Request.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"UserName": loginUserRequest.UserName}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, queryError
	}

	if !Helper.CheckPasswordHash(loginUserRequest.Password, user.Password) {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	return user, queryError
}

// CheckUserName check username exist or not
func (userRepository *UserRepository) CheckUserName(creatUserRequest Request.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}
	//queryError := Config.PostgresDB.Get(&user, `SELECT * FROM checkuserexist($1)`, creatUserRequest.UserName)

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"UserName": creatUserRequest.UserName}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, nil
	}
	return user, queryError
}

func (usserRepository *UserRepository) GetUserByUsername(username string) (Entity.User, error) {
	var user Entity.User

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"UserName": username}).Decode(&user)

	if queryError != nil {
		return Entity.User{}, queryError
	}
	return user, queryError
}

func (usserRepository *UserRepository) GetUserById(id string) (Entity.User, error) {
	var user Entity.User

	id1, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Entity.User{}, fmt.Errorf("id is not valid")
	}

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"_id": id1}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}
	return user, queryError
}

func (usserRepository *UserRepository) UpdateUser(request Request.UpdateUserRequest, userId string) (Entity.User, error) {
	var user Entity.User

	id1, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Entity.User{}, fmt.Errorf("id is not valid")
	}
	update := bson.D{
		{"$set", bson.D{{"UserName", request.UserName}, {"UpdatedAt", time.Now()}}},
	}
	result := Config.UserCollection.FindOneAndUpdate(Config.DBContext, bson.M{"_id": id1}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)

	if result != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}
	return user, result
}

func (usserRepository *UserRepository) ChangePassword(request Request.ChangePasswordRequest, userId string) (Entity.User, error) {
	var user Entity.User

	id1, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return Entity.User{}, fmt.Errorf("id is not valid")
	}

	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"_id": id1}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}

	isPasswordOk := Helper.CheckPasswordHash(request.CurrentPassword, user.Password)
	if !isPasswordOk {
		return Entity.User{}, fmt.Errorf("user not found")
	}

	hashedPassword, _ := Helper.HashPassword(request.NewPassword)
	update := bson.D{
		{"$set", bson.D{{"Password", hashedPassword}, {"UpdatedAt", time.Now()}}},
	}
	result := Config.UserCollection.FindOneAndUpdate(Config.DBContext, bson.M{"_id": id1}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	if result != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func (userRepository *UserRepository) GetAllUsers(page int, limit int) ([]Entity.User, error) {
	var userList = make([]Entity.User, 0)

	userCursor, queryError := Config.UserCollection.Find(Config.DBContext, bson.M{}, Helper.NewMongoPaginate(limit, page).GetPaginatedOpts())
	if queryError != nil {
		return nil, queryError
	}

	// decode users and append to list
	for userCursor.Next(Config.DBContext) {
		var user Entity.User
		if err := userCursor.Decode(&user); err != nil {
			log.Println(err)
		}
		userList = append(userList, user)
	}

	return userList, queryError
}

func (userRepository *UserRepository) ChangeActiveStatus(userId string) (Entity.User, error) {
	var user Entity.User

	id1, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return Entity.User{}, fmt.Errorf("id is not valid")
	}

	fmt.Println("user", user, id1)
	queryError := Config.UserCollection.FindOne(Config.DBContext, bson.M{"_id": id1}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}
	fmt.Println("user", user)

	update := bson.D{
		{"$set", bson.D{{"IsActive", !user.IsActive}, {"UpdatedAt", time.Now()}}},
	}
	result := Config.UserCollection.FindOneAndUpdate(Config.DBContext, bson.M{"_id": id1}, update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&user)
	fmt.Println("user", user)
	if result != nil {
		return Entity.User{}, fmt.Errorf("user not found")
	}

	return user, nil

}
