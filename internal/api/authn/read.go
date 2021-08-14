package authn

import (
	"github.com/matt-hoiland/architecting/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *AuthNAPI) GetCredentialsByID(id primitive.ObjectID) (*data.UserCredentials, error) {
	return nil, nil
}

func (api *AuthNAPI) GetCredentialsByAccountName(accountName string) (*data.UserCredentials, error) {
	return nil, nil
}
