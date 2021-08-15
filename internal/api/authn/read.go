package authn

import (
	"context"

	"github.com/matt-hoiland/architecting/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *AuthNAPI) GetCredentialsByID(id primitive.ObjectID) (*data.UserCredentials, error) {
	return api.GetCredentials("_id", id)
}

func (api *AuthNAPI) GetCredentialsByEmail(email string) (*data.UserCredentials, error) {
	return api.GetCredentials("email", email)
}

func (api *AuthNAPI) GetCredentials(key string, value interface{}) (*data.UserCredentials, error) {
	filter := primitive.D{primitive.E{Key: key, Value: value}}
	res := api.collection.FindOne(context.TODO(), filter)

	var creds data.UserCredentials
	err := res.Decode(&creds)
	if err != nil {
		return nil, err
	}

	return &creds, nil
}
