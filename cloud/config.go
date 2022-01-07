package cloud

import "github.com/jesko-plitt/golive/env"

type MinioConfig struct {
	URL            string
	AccessKey      string
	SecretKey      string
	BucketName     string
	BucketLocation string
}

func ProvideMinioConfig() *MinioConfig {
	return &MinioConfig{
		URL:            env.Get("MINIO_URL", "minio:9000"),
		AccessKey:      env.Get("MINIO_ACCESS_KEY", ""),
		SecretKey:      env.Get("MINIO_SECRET_KEY", ""),
		BucketName:     env.Get("MINIO_BUCKET_NAME", "default"),
		BucketLocation: env.Get("MINIO_BUCKET_LOCATION", "eu_central_1"),
	}
}

// TypesenseConfig configuration
type TypesenseConfig struct {
	Server string
	ApiKey string
}

func ProvideTypesenseConfig() *TypesenseConfig {
	return &TypesenseConfig{
		Server: env.Get("TYPESENSE_SERVER", "http://typesense:8108"),
		ApiKey: env.Get("TYPESENSE_API_KEY", ""),
	}
}

// ImageProxyConfig configuration
type ImageProxyConfig struct {
	URL        string
	Prefix     string
	Key        string
	Salt       string
	BucketName string
}

func ProvideImageproxyConfig() *ImageProxyConfig {
	return &ImageProxyConfig{
		URL:        env.Get("IMAGEPROXY_URL", "http://127.0.0.1:8080"),
		Prefix:     env.Get("IMAGEPROXY_PREFIX", ""),
		Key:        env.Get("IMAGEPROXY_KEY", ""),
		Salt:       env.Get("IMAGEPROXY_SALT", ""),
		BucketName: env.Get("IMAGEPROXY_BUCKET_NAME", "default"),
	}
}

type RedisConfig struct {
	Network  string
	Addr     string
	Username string
	Password string
	DB       int
}

func ProvideRedisConfig() *RedisConfig {
	return &RedisConfig{
		Network:  env.Get("REDIS_NETWORK", ""),
		Addr:     env.Get("REDIS_ADDR", ""),
		Username: env.Get("REDIS_USERNAME", ""),
		Password: env.Get("REDIS_PASSWORD", ""),
		DB:       env.GetInt("REDIS_DB", 0),
	}
}
