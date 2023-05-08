package etl

import (
	"context"
	"etl/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"

	"etl/models/mongodb"
)

type Extractor interface {
	GetReviews(ctx context.Context, dbName string, collectionName string) ([]mongodb.Review, error)
}

type mongoExtractor struct {
	Client       *mongo.Client
	RedisStorage *storage.RedisStorage
}

func NewMongoExtractor(mongoURI string, redisStorage *storage.RedisStorage) Extractor {
	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("faliled to connect to mongodb: %v", err)
	}

	return &mongoExtractor{
		Client:       client,
		RedisStorage: redisStorage,
	}
}

func connectToMongoDB(mongoUri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (me *mongoExtractor) GetReviews(ctx context.Context, dbName string, collectionName string) ([]mongodb.Review, error) {
	collection := me.Client.Database(dbName).Collection(collectionName)

	lastExtractedTimestamp, err := me.RedisStorage.GetLastExtractedTimestamp(ctx)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"timestamp": bson.M{"$gt": lastExtractedTimestamp}}
	findOptions := options.Find().SetLimit(100).SetSort(bson.M{"timestamp": 1})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var reviews []mongodb.Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	if len(reviews) > 0 {
		newTimestamp := reviews[len(reviews)-1].Timestamp.Unix()
		if err := me.RedisStorage.UpdateLastExtractedTimestamp(ctx, newTimestamp); err != nil {
			return nil, err
		}
	}

	return reviews, nil
}
