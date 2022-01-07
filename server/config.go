package server

import "github.com/jesko-plitt/golive/env"

type Config struct {
	Addr        string
	ProxyHeader string
	ReadTimeout int
	LogFormat   string
}

func ProvideConfig() *Config {
	return &Config{
		Addr:        env.Get("SERVER_ADDR", ":8080"),
		ProxyHeader: env.Get("SERVER_PROXY_HEADER", "X-Forwarded-For"),
		ReadTimeout: env.GetInt("SERVER_READ_TIMEOUT", 3),
		LogFormat:   env.Get("SERVER_LOG_FORMAT", "${method} ${path} [${status} - ${latency}]"),
	}
}
