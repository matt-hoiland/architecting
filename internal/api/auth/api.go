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
//
// Annotations taken from https://github.com/mongodb/mongo-go-driver/blob/v1.7.1/mongo/collection.go
type Collector interface {
	// DeleteOne executes a delete command to delete at most one document from the collection.
	//
	// The filter parameter must be a document containing query operators and can be used to select the document to be
	// deleted. It cannot be nil. If the filter does not match any documents, the operation will succeed and a DeleteResult
	// with a DeletedCount of 0 will be returned. If the filter matches multiple documents, one will be selected from the
	// matched set.
	//
	// The opts parameter can be used to specify options for the operation (see the options.DeleteOptions documentation).
	//
	// For more information about the command, see https://docs.mongodb.com/manual/reference/command/delete/.
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)

	// FindOne executes a find command and returns a SingleResult for one document in the collection.
	//
	// The filter parameter must be a document containing query operators and can be used to select the document to be
	// returned. It cannot be nil. If the filter does not match any documents, a SingleResult with an error set to
	// ErrNoDocuments will be returned. If the filter matches multiple documents, one will be selected from the matched set.
	//
	// The opts parameter can be used to specify options for this operation (see the options.FindOneOptions documentation).
	//
	// For more information about the command, see https://docs.mongodb.com/manual/reference/command/find/.
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult

	// InsertOne executes an insert command to insert a single document into the collection.
	//
	// The document parameter must be the document to be inserted. It cannot be nil. If the document does not have an _id
	// field when transformed into BSON, one will be added automatically to the marshalled document. The original document
	// will not be modified. The _id can be retrieved from the InsertedID field of the returned InsertOneResult.
	//
	// The opts parameter can be used to specify options for the operation (see the options.InsertOneOptions documentation.)
	//
	// For more information about the command, see https://docs.mongodb.com/manual/reference/command/insert/.
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	// UpdateByID executes an update command to update the document whose _id value matches the provided ID in the collection.
	// This is equivalent to running UpdateOne(ctx, bson.D{{"_id", id}}, update, opts...).
	//
	// The id parameter is the _id of the document to be updated. It cannot be nil. If the ID does not match any documents,
	// the operation will succeed and an UpdateResult with a MatchedCount of 0 will be returned.
	//
	// The update parameter must be a document containing update operators
	// (https://docs.mongodb.com/manual/reference/operator/update/) and can be used to specify the modifications to be
	// made to the selected document. It cannot be nil or empty.
	//
	// The opts parameter can be used to specify options for the operation (see the options.UpdateOptions documentation).
	//
	// For more information about the command, see https://docs.mongodb.com/manual/reference/command/update/.
	UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
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

// DeleteCredentials removes the given data from the credentials collection.
func (api *AuthAPI) DeleteCredentials(ctx context.Context, creds *data.UserCredentials) error {
	return nil
}

// DeleteCredentialsByEmail removes the most recent credentials document with the matching email.
func (api *AuthAPI) DeleteCredentialsByEmail(ctx context.Context, email string) error {
	return nil
}

// DeleteCredentialsByID removes the most recent credentials document with the matching ID.
func (api *AuthAPI) DeleteCredentialsByID(ctx context.Context, id primitive.ObjectID) error {
	return nil
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

// UpdateCredentials updates the credentials document corresponding to the same object ID.
func (api *AuthAPI) UpdateCredentials(ctx context.Context, creds *data.UserCredentials) (*data.UserCredentials, error) {
	return nil, nil
}
