package storage

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (rs *RedisStorage) MarkAsLastProcessedID(id string) error {
	_, err := rs.Client.Set(context.Background(), "etl:last_processed_id", id, 0).Result()
	return err
}

func (rs *RedisStorage) GetLastProcessedID() (primitive.ObjectID, error) {
	id, err := rs.Client.Get(context.Background(), "etl:last_processed_id").Result()
	if err == redis.Nil {
		return primitive.NilObjectID, nil // No last processed ID exists
	} else if err != nil {
		return primitive.NilObjectID, err // Some other error occurred
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, err // Failed to convert hex string to ObjectID
	}

	return objectID, nil // Successfully got the last processed ID
}
