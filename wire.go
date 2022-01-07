//+build wireinject

package golive

import (
	"io"

	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"github.com/google/wire"
	"github.com/jesko-plitt/golive/db"
	"github.com/jesko-plitt/golive/log"
	"github.com/jesko-plitt/golive/server"
)

func InitializeLogger() logging.Logger {
	wire.Build(
		log.ProvideLevel,
		wire.InterfaceValue(new(io.Writer), (io.Writer)(nil)),
		logging.New,
	)

	return &logging.DefaultLogger{}
}

func InitializeDatabase(log logging.Logger) *storage.Controller {
	wire.Build(
		// database
		db.ProvideConfig,
		db.ProvideDialector,
		db.ProvideDB,
	)

	return &storage.Controller{}
}

func InitializeServer() *Server {
	wire.Build(
		// logging
		log.ProvideLevel,
		wire.InterfaceValue(new(io.Writer), (io.Writer)(nil)),
		logging.New,

		// database
		db.ProvideConfig,
		db.ProvideDialector,
		db.ProvideDB,

		// server
		provideServer,
		server.ProvideShutdownService,
		server.ProvideConfig,
		server.ProvideRouter,
	)

	return &Server{}
}
