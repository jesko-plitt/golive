package db

import (
	"github.com/ao-concepts/logging"
	"github.com/ao-concepts/storage"
	"github.com/jesko-plitt/golive/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideConfig() *storage.Config {
	return &storage.Config{
		MaxOpenConnections: env.GetInt("DATABASE_MAX_IDLE_CONN", 0),
		MaxIdleConnections: env.GetInt("DATABASE_MAX_OPEN_CONN", 0),
	}
}

func ProvideDialector() gorm.Dialector {
	return mysql.Open(env.Require("DATABASE_DSN"))
}

func ProvideDB(dialector gorm.Dialector, config *storage.Config, log logging.Logger) *storage.Controller {
	db, err := storage.New(dialector, config, log)

	if err != nil {
		log.ErrFatal(err)
	}

	return db
}
