package Controller

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang_monolithic_bilerplate/Common/Config"
	"golang_monolithic_bilerplate/Common/Helper"
	"golang_monolithic_bilerplate/Components/User/Entity"
	"golang_monolithic_bilerplate/Components/User/Request"
	"log"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser exec query for create new user in database
func (userRepository *UserRepository) CreateUser(creatUserRequest Request.CreateUserRequest, password string) (Entity.User, error) {
	user := Entity.User{}

	result, err := Config.UserCollection.InsertOne(Config.DBCtx, Entity.User{ID: primitive.NewObjectID(),
		UserName: creatUserRequest.UserName, Password: password, CreatedAt: time.Now()})
	if err != nil {
		return Entity.User{}, err
	}

	if err = Config.UserCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		return Entity.User{}, err
	}

	return user, err
}

// LoginUser for login users
func (userRepository *UserRepository) LoginUser(loginUserRequest Request.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}

	queryError := Config.UserCollection.FindOne(Config.DBCtx, bson.M{"UserName": loginUserRequest.UserName}).Decode(&user)
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

	queryError := Config.UserCollection.FindOne(Config.DBCtx, bson.M{"UserName": creatUserRequest.UserName}).Decode(&user)
	if queryError != nil {
		return Entity.User{}, nil
	}
	return user, queryError
}

func (usserRepository UserRepository) GetUserByUsername(username string) (Entity.User, error) {
	var user Entity.User

	queryError := Config.UserCollection.FindOne(Config.DBCtx, bson.M{"UserName": username}).Decode(&user)

	if queryError != nil {
		return Entity.User{}, queryError
	}
	return user, queryError
}

func (userRepository UserRepository) GetAllUsers(page int, limit int) ([]Entity.User, error) {
	var userList = make([]Entity.User, 0)

	userCursor, queryError := Config.UserCollection.Find(Config.DBCtx, bson.M{}, Helper.NewMongoPaginate(limit, page).GetPaginatedOpts())
	log.Println("query error : ", userCursor.Err())
	if queryError != nil {
		return nil, queryError
	}

	// decode users and append to list
	for userCursor.Next(Config.DBCtx) {
		var user Entity.User
		if err := userCursor.Decode(&user); err != nil {
			log.Println(err)
		}
		userList = append(userList, user)
	}

	return userList, queryError
}
