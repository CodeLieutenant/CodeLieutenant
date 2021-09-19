package database

import (
	"time"

	"github.com/malusev998/malusev998/server/tests"
)

type Cfg struct {
	Conn string
}

func (Cfg) LazyConnect() bool {
	return true
}

func (Cfg) MaxConnLifetime() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) MaxConnIdleTime() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) HealthCheckPeriod() time.Duration {
	return 2 * time.Millisecond
}

func (Cfg) MaxConnections() int32 {
	return 5
}

func (Cfg) MinConnections() int32 {
	return 2
}

func (c Cfg) String() string {
	if c.Conn == "" {
		return tests.GetConnectionString("dusanmalusev")
	}

	return c.Conn
}
