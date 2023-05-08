package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type RedisStorage struct {
	Client *redis.Client
}

func NewRedisStorage(redisAddr string) *RedisStorage {
	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	return &RedisStorage{Client: client}
}

func (rs *RedisStorage) GetLastExtractedTimestamp(ctx context.Context) (int64, error) {
	lastExtractedTimestamp, err := rs.Client.Get(ctx, "last_extracted_timestamp").Int64()
	if err != nil && err != redis.Nil {
		return 0, err
	}

	return lastExtractedTimestamp, nil
}

func (rs *RedisStorage) UpdateLastExtractedTimestamp(ctx context.Context, timestamp int64) error {
	return rs.Client.Set(ctx, "last_extracted_timestamp", timestamp, 0).Err()
}
