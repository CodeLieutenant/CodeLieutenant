package cmd

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/malusev998/dusanmalusev/config"
	"github.com/malusev998/dusanmalusev/container"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dusanmalusev",
		Short: "Backand golang application with fiber for Dusan's website",
	}
	loggingLevel string
	cfgFile      string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&loggingLevel, "level", "debug", "Global logging level")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", ".", "Config file location")

	_ = viper.BindPFlag("global_logging_level", rootCmd.PersistentFlags().Lookup("level"))
	_ = viper.BindPFlag("config_file", rootCmd.PersistentFlags().Lookup("config"))
}

func Execute(ctx context.Context, version string, cfg *config.Config, logger zerolog.Logger) error {
	rootCmd.Version = version

	c := &container.Container{
		Ctx:    ctx,
		Logger: logger,
		Config: cfg,
	}

	rootCmd.AddCommand(newServerCommand(c))

	return rootCmd.Execute()
}
