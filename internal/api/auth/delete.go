package auth

import (
	"github.com/matt-hoiland/architecting/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *AuthAPI) DeleteCredentialsByID(id primitive.ObjectID) error {
	return nil
}

func (api *AuthAPI) DeleteCredentialsByName(accountName string) error {
	return nil
}

func (api *AuthAPI) DeleteCredentials(creds *data.UserCredentials) error {
	return nil
}
