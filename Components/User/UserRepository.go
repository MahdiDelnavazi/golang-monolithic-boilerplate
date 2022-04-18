package Controller

import (
	"fmt"
	"golang_monolithic_bilerplate/Common/Config"
	"golang_monolithic_bilerplate/Common/Helper"
	"golang_monolithic_bilerplate/Components/User/Entity"
	"golang_monolithic_bilerplate/Components/User/Request"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// CreateUser exec query for create new user in database
func (userRepository *UserRepository) CreateUser(creatUserRequest Request.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}
	password, err := Helper.HashPassword(creatUserRequest.Password)
	if err != nil {
		fmt.Println("user retun in err password : ", user)
		return Entity.User{}, err
	}
	queryError := Config.PostgresDB.Get(&user, `SELECT * FROM newuser($1 , $2, 0)`, creatUserRequest.UserName, password)
	if queryError != nil {
		fmt.Println("query error : ", queryError)

		return Entity.User{}, queryError
	}
	return user, queryError
}

// LoginUser for login users
func (userRepository *UserRepository) LoginUser(loginUserRequest Request.LoginUserRequest) (Entity.User, error) {
	user := Entity.User{}
	queryError := Config.PostgresDB.Get(&user, `SELECT * FROM loginuser($1)`, loginUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	if !Helper.CheckPasswordHash(loginUserRequest.Password, user.Password) {
		return Entity.User{}, fmt.Errorf("user or password is incorrect")
	}
	return user, queryError
}

// CheckUserName check username exist or not
func (userRepository *UserRepository) CheckUserName(creatUserRequest Request.CreateUserRequest) (Entity.User, error) {
	user := Entity.User{}

	queryError := Config.PostgresDB.Get(&user, `SELECT * FROM checkuserexist($1)`, creatUserRequest.UserName)
	fmt.Println("user after run query", queryError)
	if queryError != nil {
		return Entity.User{}, nil
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
func (userRepository *UserRepository) GetUser(creatUserRequest Request.GetUserRequest) (Entity.User, error) {
	user := Entity.User{}
	queryError := Config.PostgresDB.Get(&user, `SELECT * FROM getuser($1)`, creatUserRequest.UserName)
	if queryError != nil {
		return Entity.User{}, queryError
	}
	return user, queryError
}
