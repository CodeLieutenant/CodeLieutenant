package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type item struct {
	database string
	expected string
}

func TestGetURNConnectionString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("Env", func(t *testing.T) {
		assert.NoError(os.Setenv(DBURNEnvironmentalVariable, "postgres://postgres:postgres@localhost:5432/?sslmode=disable"))

		data := []item{
			{database: "", expected: "postgres://postgres:postgres@localhost:5432/?sslmode=disable"},
			{database: "test_database", expected: "postgres://postgres:postgres@localhost:5432/test_database?sslmode=disable"},
		}

		for _, i := range data {
			assert.Equal(i.expected, GetURNConnectionString(i.database))
		}

		assert.NoError(os.Unsetenv(DBURNEnvironmentalVariable))
	})

	t.Run("NoEnv", func(t *testing.T) {
		data := []item{
			{database: "", expected: "postgres://postgres:postgres@localhost:5432/?sslmode=disable"},
			{database: "test_database", expected: "postgres://postgres:postgres@localhost:5432/test_database?sslmode=disable"},
		}

		for _, i := range data {
			assert.Equal(i.expected, GetURNConnectionString(i.database))
		}
	})
}

func TestGetConnectionString(t *testing.T) {
	t.Parallel()
	assert := require.New(t)


	t.Run("Env", func(t *testing.T) {
		assert.NoError(os.Setenv(DBEnvironmentalVariable, "user=postgres password=postgres host=localhost port=5432"))

		data := []item{
			{database: "", expected: "user=postgres password=postgres host=localhost port=5432"},
			{database: "test_database", expected: "user=postgres password=postgres host=localhost port=5432 dbname=test_database"},
		}

		for _, i := range data {
			assert.Equal(i.expected, GetConnectionString(i.database))
		}

		assert.NoError(os.Unsetenv(DBEnvironmentalVariable))
	})

	t.Run("NoEnv", func(t *testing.T) {
		data := []item{
			{database: "", expected: "user=postgres password=postgres host=localhost port=5432"},
			{database: "test_database", expected: "user=postgres password=postgres host=localhost port=5432 dbname=test_database"},
		}

		for _, i := range data {
			assert.Equal(i.expected, GetConnectionString(i.database))
		}
	})
}
