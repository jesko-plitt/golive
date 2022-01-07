package log

import (
	"github.com/ao-concepts/logging"
	"github.com/jesko-plitt/golive/env"
)

func ProvideLevel() logging.Level {
	switch env.Get("LOG_LEVEL", "debug") {
	case "info":
		return logging.Info
	case "warn":
		return logging.Warn
	case "error":
		return logging.Error
	case "fatal":
		return logging.Fatal
	default:
		return logging.Debug
	}
}
