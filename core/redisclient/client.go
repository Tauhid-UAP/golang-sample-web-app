package redisclient

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

var PubSubClient *redis.Client

func Init() {
	Client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB: 0,
	})

	PubSubClient = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_PUB_SUB_ADDR"),
		Password: os.Getenv("REDIS_PUB_SUB_PASSWORD"),
		DB: 0,
	})
}

func PublishToRoom(ctx context.Context, room string, payload []byte) error {
	return PubSubClient.Publish(ctx, room, payload).Err()
}

func SubscribeToRoom(ctx context.Context, room string) *redis.PubSub {
	return PubSubClient.Subscribe(ctx, room)
}

func Ping(ctx context.Context) error {
	return Client.Ping(ctx).Err()
}

func PingPubSub(ctx context.Context) error {
	return PubSubClient.Ping(ctx).Err()
}
