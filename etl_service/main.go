package main

import (
	"context"
	"etl/core"
	"etl/etl"
	"etl/storage"
	"fmt"
	"log"
)

func main() {
	uri, dbName, collectionName := core.LoadMongoConfig()
	redisAddr := core.LoadRedisConfig()
	clickhouseURI := core.LoadClickhouseConfig()
	ctx := context.Background()

	redisStorage := storage.NewRedisStorage(redisAddr)
	mongoExtractor := etl.NewMongoExtractor(uri, redisStorage)
	clickhouseLoader := etl.NewClickhouseLoader(clickhouseURI)

	reviews, err := mongoExtractor.GetReviews(ctx, dbName, collectionName)
	if err != nil {
		log.Fatal("failed extract reviews: %v", err)
	}

	err = clickhouseLoader.LoadReviews(ctx, reviews)
	if err != nil {
		log.Fatalf("failed load reviews: %v", err)
	}

	fmt.Println(reviews)

}
