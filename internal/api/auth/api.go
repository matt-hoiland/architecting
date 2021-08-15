package auth

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// UserDatabase names the mongo database.
	UserDatabase = "user"

	// CredentialsCollection names the credentials collection
	CredentialsCollection = "credentials"
)

type AuthAPI struct {
	collection DBCollection
}

func NewAuthAPI(collection DBCollection) *AuthAPI {
	return &AuthAPI{
		collection: collection,
	}
}

// API Dependencies

type DBCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
}
