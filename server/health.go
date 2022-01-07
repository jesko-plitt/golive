package server

import (
	"os"
	"runtime"

	"github.com/ao-concepts/storage"
	"github.com/gofiber/fiber/v2"
)

func HealthRoute(db *storage.Controller) func(c *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sqldb, _ := db.Gorm().DB()
		stats := sqldb.Stats()

		name, _ := os.Hostname()

		return ctx.JSON(&fiber.Map{
			"db": &fiber.Map{
				"maxOpenConnections": stats.MaxOpenConnections,
				"openConnections":    stats.OpenConnections,
				"inUse":              stats.InUse,
				"idle":               stats.Idle,
				"waitCount":          stats.WaitCount,
				"waitDuration":       stats.WaitDuration,
				"maxIdleClosed":      stats.MaxIdleClosed,
				"maxIdleTimeClosed":  stats.MaxIdleTimeClosed,
				"maxLifetimeClosed":  stats.MaxLifetimeClosed,
			},
			"sys": &fiber.Map{
				"goRoutines": runtime.NumGoroutine(),
				"gomaxprocs": runtime.GOMAXPROCS(0),
				"cpus":       runtime.NumCPU(),
				"server":     name,
			},
		})
	}
}
