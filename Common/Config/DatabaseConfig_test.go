package Config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"testing"
)

type MongoDBTest struct {
	Url    string `env:"DB_URL" env-default:"mongodb://localhost:27017"`
	DBname string `env:"DB_NAME" env-default:"golang_monolithic_boilerplate"`
}

func Test_MongoDatabaseOpen(t *testing.T) {
	config := MongoDBTest{}
	err := cleanenv.ReadConfig("../../.test.env", &config)

	require.NoError(t, err)
	require.NotNil(t, config)

	client, ctx, cansel, err := connect(config.Url)
	require.NoError(t, err)
	require.NotNil(t, ctx)
	require.NotNil(t, cansel)

	dbtest := client.Database(config.DBname)
	require.NotNil(t, dbtest)

	err = client.Ping(ctx, readpref.Primary())
	require.NoError(t, err)

}
