package storage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"

	"data-processing/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

func NewMongoStorage(uri string, dbName string, collectionName string) (Storage, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoStorage{
		Client:     client,
		Collection: collection,
	}, nil
}

func (m *MongoStorage) CreateReview(review models.ReviewCreate) error {
	objectID, err := m.Collection.InsertOne(context.Background(), review)
	if err != nil {
		log.Printf("Error creating review: %v", err)
		return err
	}

	log.Printf("Review created with ID %v", objectID)
	return nil
}

func (m *MongoStorage) FindReviews() ([]models.Review, error) {
	cursor, err := m.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var reviews []models.Review
	err = cursor.All(context.Background(), &reviews)
	if err != nil {
		return nil, err
	}

	return reviews, nil
}
