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

func (api *AuthAPI) FindCredentials(key string, value interface{}) (*data.UserCredentials, error) {
	filter := primitive.D{primitive.E{Key: key, Value: value}}
	res := api.credentials.FindOne(context.TODO(), filter)

	var creds data.UserCredentials
	err := res.Decode(&creds)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}

func (api *AuthAPI) FindCredentialsByEmail(email string) (*data.UserCredentials, error) {
	return api.FindCredentials("email", email)
}

func (api *AuthAPI) FindCredentialsByID(id primitive.ObjectID) (*data.UserCredentials, error) {
	return api.FindCredentials("_id", id)
}

func (api *AuthAPI) InsertCredentials(creds *data.UserCredentials) (primitive.ObjectID, error) {
	result, err := api.credentials.InsertOne(context.TODO(), creds)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return primitive.NilObjectID, errors.New("unrecognized type returned")
}

func (api *AuthAPI) RemoveCredentials(creds *data.UserCredentials) error {
	return nil
}

func (api *AuthAPI) RemoveCredentialsByEmail(email string) error {
	return nil
}

func (api *AuthAPI) RemoveCredentialsByID(id primitive.ObjectID) error {
	return nil
}

func (api *AuthAPI) UpdateCredentials(creds *data.UserCredentials) (*data.UserCredentials, error) {
	return nil, nil
}
