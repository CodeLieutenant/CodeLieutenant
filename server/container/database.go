package container

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"

	"github.com/malusev998/malusev998/server/database"
)

func (c *Container) GetDatabasePool() *pgxpool.Pool {
	if c.DB == nil {
		var err error
		c.DB, err = database.ConnectDB(c.Ctx, database.Config{
			URL:                   c.Config.Database.URI,
			Host:                  c.Config.Database.Host,
			User:                  c.Config.Database.User,
			Password:              c.Config.Database.Password,
			DbName:                c.Config.Database.DBName,
			TimeZone:              c.Config.Database.TimeZone,
			MaxConnectionLifetime: c.Config.Database.MaxConnectionIdleTime,
			MaxConnectionIdleTime: c.Config.Database.MaxConnectionIdleTime,
			HealthCheck:           c.Config.Database.HealthCheck,
			MaxConns:              c.Config.Database.MaxConns,
			MinConns:              c.Config.Database.MinConns,
			Port:                  uint16(c.Config.Database.Port),
			SslMode:               c.Config.Database.SSLMode,
			Lazy:                  c.Config.Database.Lazy,
		}, c.Logger)

		if err != nil {
			log.Fatal().
				Err(err).
				Msg("Error while connecting to database")
			c.Logger.
				Fatal().
				Err(err).
				Msg("Error while connecting to database")
		}
	}

	return c.DB
}
