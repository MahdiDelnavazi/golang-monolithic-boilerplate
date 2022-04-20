package Config

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/url"
	"strings"
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
	Url    string
	DBname string
}

var DBMongo *mongo.Database
var DBPostgres *sqlx.DB
var err error

var (
	UserCollection   *mongo.Collection
	TicketCollection *mongo.Collection
	DBCtx            = context.TODO()
)

// DatabaseOpen Open knows how to open a database connection based on the configuration.
func DatabaseOpen(cfg DatabaseConfig) {
	sslMode := "disable"
	if cfg.DisableTLS {
		sslMode = "disable"
	}

	q := make(url.Values)
	q.Set("sslmode", sslMode)
	q.Set("timezone", "utc")
	q.Set("user", cfg.User)
	q.Set("password", cfg.Password)
	q.Set("dbname", cfg.Name)

	databaseConfig := strings.Replace(q.Encode(), "&", " ", -1)
	fmt.Println("this is database : ", databaseConfig)
	DBPostgres, err = sqlx.Open("postgres", databaseConfig)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	DBPostgres.SetMaxIdleConns(cfg.MaxIdleConns)
	DBPostgres.SetMaxOpenConns(cfg.MaxOpenConns)

}

func MongoDatabaseOpen(cfg MongoDB) {
	// Get Client, Context, CancelFunc and
	// err from connect method.
	client, ctx, cancel, err := connect(cfg.Url)
	if err != nil {
		panic(err)
	}

	DBMongo = client.Database(cfg.DBname)
	fmt.Println("this is mongo : ", client, ctx, cancel)

	// Release resource when the main
	// function is returned.
	//defer close(client, ctx, cancel)

	// Ping mongoDB with Ping method
	ping(client, ctx)
	UserCollection = DBMongo.Collection("User")
	TicketCollection = DBMongo.Collection("Ticket")
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
