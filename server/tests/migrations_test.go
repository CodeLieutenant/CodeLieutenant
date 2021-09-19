package tests

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindMigrations(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	t.Run("Env", func(t *testing.T) {
		err := os.Setenv(MigrationEnv, "TestMigrationsDir")
		assert.NoError(err)

		assert.Equal("TestMigrationsDir", findMigrations())

		assert.NoError(os.Unsetenv(MigrationEnv))
	})

	t.Run("NoEnv", func(t *testing.T) {
		dir, err := os.Getwd()
		assert.NoError(err)

		migrationsPath := path.Join(dir, "..", "database", "migrations")
		assert.Equal(migrationsPath, findMigrations())
	})
}
