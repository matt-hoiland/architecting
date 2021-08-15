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

// Collector represents the expected operations that can be performed on a mongodb collection.
// The intended implementation is mongodb's official driver in go. Interfacing allows this
// dependency to be mocked for testing.
type Collector interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
}

// AuthAPI owns and performs database transactions to the credentials collection as well as
// future authn and authz operations.
type AuthAPI struct {
	credentials Collector
}

// NewAuthAPI constructs a new AuthAPI instance with the given Collector.
func NewAuthAPI(credentials Collector) *AuthAPI {
	return &AuthAPI{
		credentials: credentials,
	}
}
