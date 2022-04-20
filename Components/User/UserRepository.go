package Controller

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang_monolithic_bilerplate/Common/Config"
	"golang_monolithic_bilerplate/Common/Helper"
	"golang_monolithic_bilerplate/Components/User/Entity"
	"golang_monolithic_bilerplate/Components/User/Request"
	"time"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser exec query for create new user in database
func (userRepository *UserRepository) CreateUser(creatUserRequest Request.CreateUserRequest, password string) (Entity.UserMongo, error) {
	user := Entity.UserMongo{}

	result, err := Config.UserCollection.InsertOne(Config.DBCtx, Entity.UserMongo{ID: primitive.NewObjectID(),
		UserName: creatUserRequest.UserName, Password: password, CreatedAt: time.Now()})
	if err != nil {
		return Entity.UserMongo{}, err
	}

	if err = Config.UserCollection.FindOne(Config.DBCtx, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		return Entity.UserMongo{}, err
	}

	return user, err
}

// LoginUser for login users
func (userRepository *UserRepository) LoginUser(loginUserRequest Request.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}
	//queryError := Config.DB.Get(&user, `SELECT * FROM loginuser($1)`, loginUserRequest.UserName)
	//if queryError != nil {
	//	return Entity.User{}, fmt.Errorf("user or password is incorrect")
	//}

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
func (userRepository *UserRepository) CheckUserName(creatUserRequest Request.CreateUserRequest) (Entity.UserMongo, error) {
	user := Entity.UserMongo{}
	//queryError := Config.PostgresDB.Get(&user, `SELECT * FROM checkuserexist($1)`, creatUserRequest.UserName)

	queryError := Config.UserCollection.FindOne(Config.DBCtx, bson.M{"UserName": creatUserRequest.UserName}).Decode(&user)
	if queryError != nil {
		return Entity.UserMongo{}, nil
	}
	return user, queryError
}

//func (userRepository *UserRepository) LogOut(logoutReq Request.LogoutRequest, payload *token.Payload) error {
//
//	err := userRepository.redis.Set(payload.Username, logoutReq.Token, 0).Err()
//
//	if err != nil {
//		return nil
//	}
//	return err
//}

// to tikcket servesam ino call mikonam az user service
// CheckUserName check username exist or not
func (userRepository *UserRepository) GetUser(creatUserRequest Request.GetUserRequest) (Entity.UserMongo, error) {
	user := Entity.UserMongo{}
	//queryError := Config.DB.Get(&user, `SELECT * FROM getuser($1)`, creatUserRequest.UserName)
	queryError := Config.UserCollection.FindOne(Config.DBCtx, bson.M{"UserName": creatUserRequest.UserName}).Decode(&user)

	if queryError != nil {
		return Entity.UserMongo{}, queryError
	}
	return user, queryError
}
