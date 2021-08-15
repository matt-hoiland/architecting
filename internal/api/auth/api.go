package auth

import (
	"context"
	"errors"

	"github.com/matt-hoiland/architecting/data"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// FindCredentials locates the most recent single credentials document which has the given value
// for the given key.
func (api *AuthAPI) FindCredentials(ctx context.Context, key string, value interface{}) (*data.UserCredentials, error) {
	filter := primitive.D{primitive.E{Key: key, Value: value}}
	res := api.credentials.FindOne(ctx, filter)

	var creds data.UserCredentials
	err := res.Decode(&creds)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}

// FindCredentialsByEmail locates the most recent credentials document which has the given email,
func (api *AuthAPI) FindCredentialsByEmail(ctx context.Context, email string) (*data.UserCredentials, error) {
	return api.FindCredentials(ctx, "email", email)
}

// FindCredentialsByID locates the most recent credentials document which has the given ID,
func (api *AuthAPI) FindCredentialsByID(ctx context.Context, id primitive.ObjectID) (*data.UserCredentials, error) {
	return api.FindCredentials(ctx, "_id", id)
}

// InsertCredetials inserts the given data into the credentials collection and returns its object ID.
func (api *AuthAPI) InsertCredentials(ctx context.Context, creds *data.UserCredentials) (primitive.ObjectID, error) {
	result, err := api.credentials.InsertOne(ctx, creds)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return primitive.NilObjectID, errors.New("unrecognized type returned")
}

// RemoveCredentials removes the given data from the credentials collection.
func (api *AuthAPI) RemoveCredentials(ctx context.Context, creds *data.UserCredentials) error {
	return nil
}

// RemoveCredentialsByEmail removes the most recent credentials document with the matching email.
func (api *AuthAPI) RemoveCredentialsByEmail(ctx context.Context, email string) error {
	return nil
}

// RemoveCredentialsByID removes the most recent credentials document with the matching ID.
func (api *AuthAPI) RemoveCredentialsByID(ctx context.Context, id primitive.ObjectID) error {
	return nil
}

// UpdateCredentials updates the credentials document corresponding to the same object ID.
func (api *AuthAPI) UpdateCredentials(ctx context.Context, creds *data.UserCredentials) (*data.UserCredentials, error) {
	return nil, nil
}
