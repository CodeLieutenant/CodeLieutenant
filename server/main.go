package main

import (
	"context"
	"flag"
	"io"
	"os"
	"os/signal"
	"path/filepath"

	"github.com/go-playground/locales/en"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/rs/zerolog/log"

	"github.com/malusev998/dusanmalusev/api"
	"github.com/malusev998/dusanmalusev/api/routes"
	"github.com/malusev998/dusanmalusev/config"
	"github.com/malusev998/dusanmalusev/container"
	"github.com/malusev998/dusanmalusev/database"
	"github.com/malusev998/dusanmalusev/handlers"
	"github.com/malusev998/dusanmalusev/logging"
	"github.com/malusev998/dusanmalusev/validators"
)

func createLogFile(path string) (file io.WriteCloser, err error) {
	if !filepath.IsAbs(path) {
		path, err = filepath.Abs(path)
		if err != nil {
			return nil, err
		}
	}

	directory := filepath.Dir(path)
	if err := os.MkdirAll(directory, 0744); err != nil {
		return nil, err
	}

	file, err = os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)

	if err != nil {
		return nil, err
	}

	return
}

func main() {
	loggingLevel := flag.String("logging", "debug", "Global logging level")
	flag.Parse()

	logging.DefaultLogger(*loggingLevel)

	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	cfg, err := config.New("config", ".")

	if err != nil {
		log.Fatal().Err(err).Msg("Error while loading configuration")
	}

	logFile, err := createLogFile(cfg.Logging.File)

	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening log file")
	}

	dbLogFile, err := createLogFile(cfg.Database.LogFile)

	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening database log file")
	}

	logger := logging.New(
		cfg.Logging.Level,
		cfg.Logging.Console,
		logFile,
	)

	english := en.New()
	uni := ut.New(english, english)

	trans, _ := uni.GetTranslator(cfg.Locale)
	validate := validator.New()

	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		logger.Fatal().Err(err).Msg("Error while registering english translations")
	}

	if err := validators.Register(&logger, validate, trans); err != nil {
		logger.Fatal().Err(err).Msg("Error while registering custom validators")
	}

	db, err := database.ConnectDB(database.Config{
		Host:     cfg.Database.Host,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DbName:   cfg.Database.DBName,
		Port:     uint16(cfg.Database.Port),
		SslMode:  cfg.Database.SSLMode,
		TimeZone: cfg.Database.TimeZone,
	}, dbLogFile, cfg.Debug)

	if err != nil {
		log.Fatal().Err(err).Msg("Error while connecting to database")
	}

	diContainer := container.Container{
		Ctx:       ctx,
		Logger:    &logger,
		DB:        db,
		Validator: validate,
	}

	go func(cancel *context.CancelFunc) {
		s := <-signalCh
		logger.Info().Msgf("Shutting down... Signal: %s\n", s)
		(*cancel)()
	}(&cancel)

	logger.Debug().Msg("Starting HTTP Api")

	provider := api.NewFiberAPI(
		cfg.HTTP.Address,
		cfg.HTTP.Prefork,
		cfg.Debug,
		handlers.Error(diContainer.Logger, trans),
		routes.RegisterRouter,
	)

	go func() {
		<-ctx.Done()

		if err := provider.Close(); err != nil {
			logger.Error().
				Err(err).
				Msg("Error while shutting down application\n")
		}

		if err := database.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing sql database file")
		}

		if err := logFile.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing log file")
		}

		if err := dbLogFile.Close(); err != nil {
			diContainer.Logger.Err(err).Msg("Error while closing database log file")
		}

	}()

	if err := provider.Register(&diContainer); err != nil {
		logger.Fatal().
			Err(err).
			Msg("Error while configuring http server")
	}

	if err := provider.Listen(); err != nil {
		log.Fatal().Err(err).Msg("Error while starting the server")
	}
}
