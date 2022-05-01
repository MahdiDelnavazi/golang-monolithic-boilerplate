package Test

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/mahdidl/golang_boilerplate/Common/Config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type RedisConfig struct {
	Host     string
	Password string
	Port     string
}

type MongoDBTest struct {
	Url    string `env:"DB_URL" env-default:"mongodb://localhost:27017"`
	DBname string `env:"DB_NAME_TEST" env-default:"golang_monolithic_boilerplate_test"`
}

func OpenTestingDatabase() {
	config := MongoDBTest{}

	if parseError := cleanenv.ReadConfig("../../.test.env", &config); parseError != nil {
		fmt.Errorf("parsing config: %w", parseError)
	}

	fmt.Println(config)
	client, ctx, cancel, err := connect(config.Url)
	if err != nil {
		panic(err)
	}

	Config.DBMongo = client.Database(config.DBname)
	fmt.Println("this is mongo : ", client, ctx, cancel)

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("connected successfully")

	Config.UserCollection = Config.DBMongo.Collection("User")
	Config.TicketCollection = Config.DBMongo.Collection("Ticket")
	Config.RoleCollection = Config.DBMongo.Collection("Role")
	Config.PermissionCollection = Config.DBMongo.Collection("Permission")
	Config.RolePermissionCollection = Config.DBMongo.Collection("RolePermission")
}

//func OpenTestingRedis() {
//
//	config := Config.RedisConfig{}
//	if parseError := cleanenv.ReadConfig(".test.env", &config); parseError != nil {
//		fmt.Errorf("parsing config: %w", parseError)
//	}
//
//	client := redis.NewClient(&redis.Options{
//		Addr:     config.Host,
//		Password: config.Password,
//		DB:       0,
//	})
//
//	pong, err := client.Ping().Result()
//	if err != nil {
//		log.Fatalf("cannot connect to redis : %s ", err)
//	}
//	Redis = client
//	fmt.Println(pong, err)
//}

func connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// mongo.Connect return mongo.Client method
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

// This is a user defined method that accepts
// mongo.Client and context.Context
// This method used to ping the mongoDB, return error if any.
func ping(client *mongo.Client, ctx context.Context) error {

	// mongo.Client has Ping to ping mongoDB, deadline of
	// the Ping method will be determined by cxt
	// Ping method return error if any occurred, then
	// the error can be handled.
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
