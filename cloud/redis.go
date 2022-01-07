package cloud

import (
	"context"
	"encoding/json"

	"github.com/ao-concepts/logging"
	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

type RedisClient struct {
	client *redis.Client
	log    logging.Logger
	ctx    context.Context
	cancel context.CancelFunc
}

func ProvideRedisClient(log logging.Logger) *RedisClient {
	cfg := ProvideRedisConfig()
	ctx, cancel := context.WithCancel(context.Background())

	return &RedisClient{
		client: redis.NewClient(&redis.Options{
			Network:  cfg.Network,
			Addr:     cfg.Addr,
			Password: cfg.Password,
			Username: cfg.Username,
			DB:       cfg.DB,
		}),
		log:    log,
		ctx:    ctx,
		cancel: cancel,
	}
}

// EventPayload of a redis event
type EventPayload map[string]interface{}

// Bind the payload to a struct
func (p *EventPayload) Bind(s interface{}) error {
	return mapstructure.Decode(p, s)
}

// HandlerFunc is to be handled by a redis subscription
type HandlerFunc func(payload EventPayload)

// Subscribe subscribes a listener to a channel on the redis client
func (rc *RedisClient) Subscribe(channel string, handler HandlerFunc) {
	subscription := rc.client.PSubscribe(context.Background(), channel)

	go func() {
		for {
			select {
			case msg := <-subscription.Channel():
				go func() {
					defer func() {
						if r := recover(); r != nil {
							rc.log.Error("redis: recovering subscription from panic: %v", r)
						}
					}()

					var payload EventPayload

					if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
						rc.log.ErrError(err)
						return
					}

					handler(payload)
				}()
			case <-rc.ctx.Done():
				subscription.Close()
				return
			}
		}
	}()
}

// Publish a message to redis pub/sub
func (rc *RedisClient) Publish(channel string, data EventPayload) error {
	msgData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return rc.client.Publish(context.Background(), channel, msgData).Err()
}

// Shutdown the client gracefully
func (rc *RedisClient) Shutdown() error {
	rc.log.Info("redis: shutting down client")
	defer rc.log.Info("redis: client stopped")
	rc.cancel()

	return rc.client.Close()
}
