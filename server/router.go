package server

import (
	"time"

	"github.com/ao-concepts/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func ProvideRouter(cfg *Config, log logging.Logger) *fiber.App {
	router := fiber.New(fiber.Config{
		DisableStartupMessage: log.GetLevel() != logging.Debug,
		ReadTimeout:           time.Duration(cfg.ReadTimeout) * time.Second,
		ProxyHeader:           cfg.ProxyHeader,
	})

	router.Use(logger.New(logger.Config{
		Output: log,
		Format: cfg.LogFormat,
	}))

	return router
}
