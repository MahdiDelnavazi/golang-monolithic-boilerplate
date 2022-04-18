package Config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/url"
	"strings"
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

var PostgresDB *sqlx.DB
var err error

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
	PostgresDB, err = sqlx.Open("postgres", databaseConfig)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	PostgresDB.SetMaxIdleConns(cfg.MaxIdleConns)
	PostgresDB.SetMaxOpenConns(cfg.MaxOpenConns)

}
