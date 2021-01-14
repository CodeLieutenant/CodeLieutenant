package cmd

import (
	"github.com/spf13/cobra"

	"github.com/malusev998/malusev998/api"
	"github.com/malusev998/malusev998/api/routes"
	"github.com/malusev998/malusev998/container"
	"github.com/malusev998/malusev998/handlers"
)

type server struct {
	container *container.Container
	address   string
	debug     bool
	prefork   bool
}

func newServerCommand(c *container.Container) *cobra.Command {
	s := server{
		container: c,
	}

	command := &cobra.Command{
		Use:   "server",
		Short: "Start the http server serving Dusan's Website",
		RunE:  s.Execute,
	}

	command.Flags().BoolVarP(&s.debug, "debug", "d", false, "Run server in debug mode")
	command.Flags().BoolVarP(&s.prefork, "prefork", "p", false, "Run server with prefork")
	command.Flags().StringVarP(&s.address, "address", "a", "", "Address on which server will run")

	return command
}

func (s *server) Execute(cmd *cobra.Command, args []string) error {
	prefork := s.container.Config.HTTP.Prefork
	address := s.container.Config.HTTP.Address

	if s.prefork {
		prefork = true
	}

	if s.address != "" {
		address = s.address
	}

	provider := api.NewFiberAPI(
		address,
		prefork,
		s.debug,
		handlers.Error(s.container.Logger, s.container.GetTranslator()),
		routes.RegisterRouter,
	)

	go s.Close(provider)

	if err := provider.Register(s.container); err != nil {
		s.container.Logger.Fatal().
			Err(err).
			Msg("Error while configuring http server")
	}

	if err := provider.Listen(); err != nil {
		s.container.Logger.Fatal().Err(err).Msg("Error while starting the server")
	}

	return nil
}

func (s *server) Close(provider api.Interface) {
	<-s.container.Ctx.Done()
	if err := provider.Close(); err != nil {
		s.container.Logger.Error().
			Err(err).
			Msg("Error while shutting down application\n")
	}
}
