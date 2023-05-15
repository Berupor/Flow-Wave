package etl

import (
	"context"
	"database/sql"
	"etl/models/mongodb"

	"fmt"
	"github.com/ClickHouse/clickhouse-go"
	_ "github.com/ClickHouse/clickhouse-go"
	"log"
)

type Loader interface {
	LoadReviews(ctx context.Context, reviews []mongodb.Review) error
}

type clickhouseLoader struct {
	db *sql.DB
}

func NewClickhouseLoader(clickhouseURI string) Loader {
	db, err := sql.Open("clickhouse", clickhouseURI)
	if err != nil {
		log.Fatalf("failed to connect to clickhouse: %v", err)
	}

	if err := db.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			fmt.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			fmt.Println(err)
		}
	}

	return &clickhouseLoader{
		db: db,
	}
}

func (cl *clickhouseLoader) LoadReviews(ctx context.Context, reviews []mongodb.Review) error {
	tx, err := cl.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO reviews (ProductID, PlaceID, AuthorID, Rating, Review, Timestamp) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	for _, review := range reviews {
		if _, err := stmt.Exec(
			review.ProductID,
			review.PlaceID,
			review.AuthorID,
			review.Rating,
			review.Review,
			review.Timestamp,
		); err != nil {
			return err
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
