package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func URI(host, port string) string {
	return fmt.Sprintf("mongodb://%s:%s", host, port)
}

func Connect(ctx context.Context, uri, db, col string) (*mongo.Client, *mongo.Collection, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, fmt.Errorf("unable to connect to mongodb: %w", err)
	}

	collection := client.Database(db).Collection(col)

	return client, collection, nil
}
