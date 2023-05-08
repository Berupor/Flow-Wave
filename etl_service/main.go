package main

import (
	"context"
	"etl/core"
	"etl/etl"
	"etl/storage"
	"fmt"
)

func main() {
	uri, dbName, collectionName := core.LoadMongoConfig()
	redisAddr := core.LoadRedisConfig()
	ctx := context.Background()

	redisStorage := storage.NewRedisStorage(redisAddr)
	mongoExtractor := etl.NewMongoExtractor(uri, redisStorage)

	reviews, _ := mongoExtractor.GetReviews(ctx, dbName, collectionName)

	fmt.Println(reviews)

}
