package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ilyakaznacheev/cleanenv"
	"golang_monolithic_bilerplate/Common/Config"
	token "golang_monolithic_bilerplate/Common/Token"
	"golang_monolithic_bilerplate/Router"
)

func main() {

	// =====================================================
	// Open Config file
	config := Config.EnvironmentConfig{}
	if parseError := cleanenv.ReadConfig(".env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}
	fmt.Printf("%+v\n", config)

	// =====================================================
	// Config logger
	Config.NewLogger("Polaris Storage Service", config.DB.Host)

	// =====================================================
	// Open Database Connection
	//Config.MongoDatabaseOpen(Config.DatabaseConfig{
	//	User:         config.DB.User,
	//	Password:     config.DB.Password,
	//	Host:         config.DB.Host,
	//	Name:         config.DB.Name,
	//	MaxIdleConns: config.DB.MaxIdleConns,
	//	MaxOpenConns: config.DB.MaxOpenConns,
	//	DisableTLS:   config.DB.DisableTLS,
	//})
	fmt.Println("this is mongo cofig :", config.MongoDB)
	Config.MongoDatabaseOpen(Config.MongoDB{Url: config.MongoDB.Url, DBname: config.MongoDB.DBname})
	//defer func() {
	//	Config.DB.Close()
	//	fmt.Println("database in disconnected ")
	//}()

	// =====================================================
	// Create new token maker
	token.NewPasetoMaker(config.Token.TokenSymmetricKey)

	// =====================================================
	// Create new token maker
	Config.RedisConnection(config)

	// =====================================================
	// Run server
	app := gin.Default()
	app.MaxMultipartMemory = 8 << 20
	app.Static("/assets/", "./public")
	Router.Routes(app)

	errorChannel := make(chan error)
	func() {
		Config.Logger.Infow("Project Running On PORT", config.Api.ApiHost)
		errorChannel <- app.Run(config.Api.ApiHost)
	}()

}
