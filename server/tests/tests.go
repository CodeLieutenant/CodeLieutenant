package tests

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	ospath "path"
	"path/filepath"
	"time"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/malusev998/malusev998/server/validators"
)

const (
	RootProjectDir = "server"
	MigrationEnv   = "MIGRATIONS_DIR"
)

type DropDatabase func()

func GetConnectionString(database string) string {
	connectionString := os.Getenv("DB_CONN")

	if connectionString != "" {
		if database != "" {
			config, err := pgxpool.ParseConfig(connectionString)
			if err != nil {
				return ""
			}

			config.ConnConfig.Database = database
			return config.ConnString()
		}

		return connectionString
	}

	if database == "" {
		return fmt.Sprintf(
			"user=postgres password=postgres host=localhost port=5432 dbname=%s",
			database,
		)
	}

	return "user=postgres password=postgres host=localhost port=5432"
}

func GetURNConnectionString(database string) string {
	connectionString := os.Getenv("DB_CONN")

	if connectionString != "" {
		if database != "" {
			config, err := pgxpool.ParseConfig(connectionString)
			if err != nil {
				return ""
			}

			config.ConnConfig.Database = database
			return config.ConnString()
		}

		return connectionString
	}

	return fmt.Sprintf(
		"postgres://postgres:postgres@localhost:5432/%s?sslmode=disable",
		database,
	)
}

func GetValidator() *validator.Validate {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	_ = en_translations.RegisterDefaultTranslations(validate, trans)
	_ = validators.Register(validate, trans)

	return validate
}

func createDatabaseName() string {
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	i := random.Int31()

	return fmt.Sprintf("dusanmalusev_%d", i)
}

func findMigrations() string {
	env := os.Getenv("MIGRATIONS_DIR")

	if env != "" {
		return env
	}

	path, err := os.Getwd()

	if err != nil {
		panic("migration dir not found")
	}

	dir := filepath.Base(path)

	for dir != RootProjectDir {
		path = filepath.Dir(path)
		dir = filepath.Base(path)
	}

	return ospath.Join(path, "database", "migrations")
}

func CreateDatabase() (*pgxpool.Pool, DropDatabase, error) {
	migrationDir := findMigrations()

	databaseName := createDatabaseName()

	config, err := pgxpool.ParseConfig(GetURNConnectionString(""))

	if err != nil {
		return nil, nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		return nil, nil, err
	}

	_, err = pool.Exec(context.Background(), fmt.Sprintf("CREATE DATABASE %s;", databaseName))

	if err != nil {
		return nil, nil, err
	}

	drop := func(pool *pgxpool.Pool) {
		_, err = pool.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %s;", databaseName))

		if err != nil {
			panic(err)
		}
	}

	migrations := fmt.Sprintf("file://%s", migrationDir)

	m, err := migrate.New(
		migrations,
		GetURNConnectionString(databaseName),
	)

	if err != nil {
		drop(pool)
		return nil, nil, err
	}

	if err := m.Up(); err != nil {
		drop(pool)
		return nil, nil, err
	}

	config, err = pgxpool.ParseConfig(GetURNConnectionString(databaseName))

	if err != nil {
		drop(pool)
		return nil, nil, err
	}

	_, _ = m.Close()
	pool.Close()

	pool, err = pgxpool.ConnectConfig(context.Background(), config)

	if err != nil {
		drop(pool)
		return nil, nil, err
	}

	return pool, func() {
		pool.Close()

		config, err := pgxpool.ParseConfig(GetConnectionString(""))

		if err != nil {
			panic(err)
		}

		pool, err = pgxpool.ConnectConfig(context.Background(), config)

		if err != nil {
			panic(err)
		}

		defer pool.Close()

		_, err = pool.Exec(context.Background(), fmt.Sprintf("DROP DATABASE %s;", databaseName))

		if err != nil {
			panic(err)
		}
	}, nil
}
