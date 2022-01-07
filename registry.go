package golive

import (
	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"
)

type Container interface {
	Log() logging.Logger
	DB() *storage.Controller
	GetServer() *Server
	GetRouter() *fiber.App
}

type Registry struct {
	server      *Server
	log         logging.Logger
	database    *storage.Controller
	router      *fiber.App
	minioClient *minio.Client
}

func Init() *Registry {
	srv := InitializeServer()

	return &Registry{
		server:   srv,
		log:      srv.Logger(),
		database: srv.DB(),
		router:   srv.Router(),
	}
}

func (r *Registry) GetServer() *Server {
	if r.server == nil {
		r.Log().Fatal("Container: No server service")
	}

	return r.server
}

func (r *Registry) Log() logging.Logger {
	if r.log == nil {
		panic("Container: No logging service")
	}

	return r.log
}

func (r *Registry) DB() *storage.Controller {
	if r.database == nil {
		r.Log().Fatal("Container: No database service")
	}

	return r.database
}

func (r *Registry) GetRouter() *fiber.App {
	if r.router == nil {
		r.Log().Fatal("Container: No router service")
	}

	return r.router
}
