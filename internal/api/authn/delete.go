package authn

import (
	"github.com/matt-hoiland/architecting/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *AuthNAPI) DeleteCredentialsByID(id primitive.ObjectID) error {
	return nil
}

func (api *AuthNAPI) DeleteCredentialsByName(accountName string) error {
	return nil
}

func (api *AuthNAPI) DeleteCredentials(creds *data.UserCredentials) error {
	return nil
}
