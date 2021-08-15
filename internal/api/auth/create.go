package auth

import (
	"context"
	"errors"

	"github.com/matt-hoiland/architecting/data"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (api *AuthAPI) InsertCredentials(creds *data.UserCredentials) (primitive.ObjectID, error) {
	result, err := api.collection.InsertOne(context.TODO(), creds)
	if err != nil {
		log.Error(err)
		return primitive.NilObjectID, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return primitive.NilObjectID, errors.New("unrecognized type returned")
}
