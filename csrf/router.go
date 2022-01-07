package csrf

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/storage/redis"
)

func UseRedis(router *fiber.App) {
	cfg := ProvideRedisConfig()

	router.Use(csrf.New(csrf.Config{
		CookieName:     cfg.CookieName,
		CookieSameSite: cfg.CookieSameSite,
		Storage: redis.New(redis.Config{
			Host:     cfg.Host,
			Port:     cfg.Port,
			Username: cfg.Username,
			Password: cfg.Password,
			Database: cfg.DB,
			Reset:    false,
		}),
	}))
}
