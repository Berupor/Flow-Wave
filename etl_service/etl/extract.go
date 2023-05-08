package etl

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"etl/models/mongodb"
)

type Extractor interface {
	GetReviews() ([]mongodb.Review, error)
}

type mongoExtractor struct {
	client *mongo.Client
}

func NewMongoExtractor(mongoURI string) Extractor {
	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("faliled to connect to mongodb: %v", err)
	}

	return &mongoExtractor{
		client: client,
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

func (me *mongoExtractor) GetReviews() ([]mongodb.Review, error) {
	collection := me.client.Database("database_name").Collection("reviews")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var reviews []mongodb.Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}

	return reviews, nil
}
