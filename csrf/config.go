package csrf

import "github.com/jesko-plitt/golive/env"

type RedisConfig struct {
	CookieName     string
	CookieSameSite string
	Host           string
	Port           int
	Username       string
	Password       string
	DB             int
}

func ProvideRedisConfig() *RedisConfig {
	return &RedisConfig{
		CookieName:     env.Get("CSRF_COOKIE_NAME", "csrf"),
		CookieSameSite: env.Get("CSRF_COOKIE_SAME_SITE", "Lax"),
		Host:           env.Get("CSRF_REDIS_HOST", "127.0.0.1"),
		Port:           env.GetInt("CSRF_REDIS_PORT", 6379),
		Username:       env.Get("CSRF_REDIS_USERNAME", ""),
		Password:       env.Get("CSRF_REDIS_PASSWORD", ""),
		DB:             env.GetInt("CSRF_REDIS_DB", 0),
	}
}
