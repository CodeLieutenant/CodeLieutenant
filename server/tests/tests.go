package tests

import (
	"context"
	"os"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/suite"

	"github.com/malusev998/malusev998/validators"
)

type DBTestCase struct {
	Ctx              context.Context
	DB              *pgxpool.Pool
	ConnectionString string
	suite.Suite
}

func (s *DBTestCase) SetupSuite() {
	config, err := pgxpool.ParseConfig(s.ConnectionString)
	if err != nil {
		s.FailNow("Failed to create pgxpool config", err.Error())
	}

	s.DB, err = pgxpool.ConnectConfig(s.Ctx, config)

	if err != nil {
		s.FailNow("Failed to connect to pgxpool", err.Error())
	}
}

func (s *DBTestCase) TearDownSuite() {
	s.DB.Close()
}

func GetConnectionString() string {
	connectionString := os.Getenv("DB_CONN")

	if connectionString != "" {
		return connectionString
	}

	return "postgres://postgres:postgres@localhost:5432/dusanmalusev?sslmode=disable"
}



func GetValidator() *validator.Validate {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	validate := validator.New()

	en_translations.RegisterDefaultTranslations(validate, trans)
	validators.Register(validate, trans)

	return validate
}
