package Config

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

// DatabaseConfig Config is the required properties to use the database.
type DatabaseConfig struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

type MongoDB struct {
	Url    string `env:"DB_URL" env-default:"mongodb://localhost:27017"`
	DBname string `env:"DB_NAME" env-default:"golang_monolithic_boilerplate"`
}

var (
	DBMongo                  *mongo.Database
	UserCollection           *mongo.Collection
	TicketCollection         *mongo.Collection
	RoleCollection           *mongo.Collection
	PermissionCollection     *mongo.Collection
	RolePermissionCollection *mongo.Collection
	IngredientCollection     *mongo.Collection
	DBCtx                    = context.TODO()
	err                      error
)

func MongoDatabaseOpen(cfg MongoDB) {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, _, err := connect(cfg.Url)
	if err != nil {
		panic(err)
	}

	DBMongo = client.Database(cfg.DBname)

	ping(client, ctx)
	UserCollection = DBMongo.Collection("User")
	TicketCollection = DBMongo.Collection("Ticket")
	RoleCollection = DBMongo.Collection("Role")
	PermissionCollection = DBMongo.Collection("Permission")
	RolePermissionCollection = DBMongo.Collection("RolePermission")
	IngredientCollection = DBMongo.Collection("Ingredient")
}

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
