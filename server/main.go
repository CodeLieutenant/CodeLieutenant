package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"

	"github.com/malusev998/dusanmalusev/cmd"
	"github.com/malusev998/dusanmalusev/config"
	"github.com/malusev998/dusanmalusev/logging"
	"github.com/malusev998/dusanmalusev/utils"
)

const (
	Version = "dev"
)

func main() {
	logging.DefaultLogger(viper.GetString("global_logging_level"))

	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	cfg, err := config.New("config", viper.GetString("config_file"))
	if err != nil {
		log.Fatal().Err(err).Msg("Error while loading configuration")
	}

	logFile, err := utils.CreateFile(cfg.Logging.File)
	if err != nil {
		log.Fatal().Err(err).Msg("Error while opening log file")
	}

	logger := logging.New(
		cfg.Logging.Level,
		cfg.Logging.Console,
		logFile,
	)

	go func(cancel *context.CancelFunc) {
		s := <-signalCh
		logger.Info().Msgf("Shutting down... Signal: %s\n", s)
		(*cancel)()
	}(&cancel)

	go func() {
		<-ctx.Done()

		if err := logFile.Close(); err != nil {
			logger.Err(err).Msg("Error while closing log file")
		}
	}()

	if err := cmd.Execute(ctx, Version, &cfg, logger); err != nil {
		os.Exit(1)
	}
}
