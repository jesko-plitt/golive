package golive

import (
	"sync"

	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/jesko-plitt/golive/server"
)

type Server struct {
	cfg      *server.Config
	shutdown *server.ShutdownService
	log      logging.Logger
	db       *storage.Controller
	router   *fiber.App
}

func (s *Server) Logger() logging.Logger {
	return s.log
}

func (s *Server) DB() *storage.Controller {
	return s.db
}

func (s *Server) Shutdown() *server.ShutdownService {
	return s.shutdown
}

func (s *Server) Serve() {
	wg := &sync.WaitGroup{}
	s.shutdown.ListenForShutdown(s.router, wg)

	wg.Add(1)
	if err := s.router.Listen(s.cfg.Addr); err != nil {
		s.log.ErrFatal(err)
	}
	wg.Wait()
}

func (s *Server) Router() *fiber.App {
	return s.router
}

func provideServer(
	cfg *server.Config,
	router *fiber.App,
	shutdown *server.ShutdownService,
	log logging.Logger,
	db *storage.Controller,
) *Server {
	router.Get("/health", server.HealthRoute(db))

	return &Server{
		cfg:      cfg,
		shutdown: shutdown,
		router:   router,
		log:      log,
		db:       db,
	}
}

func ProvideLoggerFromServer(srv *Server) logging.Logger {
	return srv.log
}

func ProvideDatabaseFromServer(srv *Server) *storage.Controller {
	return srv.db
}

func ProvideShutdownFromServer(srv *Server) *server.ShutdownService {
	return srv.shutdown
}
