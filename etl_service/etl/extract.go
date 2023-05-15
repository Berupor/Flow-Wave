package etl

import (
	"context"
	"etl/models/mongodb"
	"etl/storage"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Extractor interface {
	GetReviews(ctx context.Context, dbName string, collectionName string) ([]mongodb.Review, error)
}

type mongoExtractor struct {
	Client       *mongo.Client
	RedisStorage *storage.RedisStorage
}

func NewMongoExtractor(mongoURI string, storage *storage.RedisStorage) Extractor {
	client, err := connectToMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("failed to connect to mongodb: %v", err)
	}

	return &mongoExtractor{
		Client:       client,
		RedisStorage: storage,
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

	// Get the last processed _id from Redis
	lastID, err := me.RedisStorage.GetLastProcessedID()
	if err != nil {
		return nil, err
	}

	filter := bson.M{}
	if lastID != primitive.NilObjectID {
		// Only get reviews with an _id greater than the last processed _id.
		filter["_id"] = bson.M{"$gt": lastID}
	}

	findOptions := options.Find().SetLimit(100).SetSort(bson.D{{"_id", 1}})

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var reviews []mongodb.Review
	for cursor.Next(ctx) {
		var review mongodb.Review
		if err = cursor.Decode(&review); err != nil {
			return nil, err
		}

		// Only mark as the last processed after it's been successfully processed
		if err = me.RedisStorage.MarkAsLastProcessedID(review.ID.Hex()); err != nil {
			return nil, err
		}

		reviews = append(reviews, review)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return reviews, nil
}
