package data

import (
	"encoding/hex"
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserCredentials represents the document model stored in mongo.
type UserCredentials struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"        json:"_id,omitempty"`
	Email                string             `bson:"email"                json:"email"`
	Hash                 ByteSlice          `bson:"hash"                 json:"-"`
	Salt                 ByteSlice          `bson:"salt"                 json:"-"`
	Validated            bool               `bson:"validated"            json:"validated"`
	TimestampUpdatedLast time.Time          `bson:"ts_updated,omitempty" json:"ts_updated,omitempty"`
}

type ByteSlice []byte

func (slice *ByteSlice) UnmarshalBSONValue(btype bsontype.Type, data []byte) error {
	if btype != bsontype.String {
		return errors.New("expected bson string, received " + btype.String())
	}
	var val string
	err := bson.Unmarshal(data, &val)
	if err != nil {
		log.Error(err)
		return err
	}
	*slice = []byte(val)
	return nil
}

func (slice *ByteSlice) MarshalBSONValue() (bsontype.Type, []byte, error) {
	val := hex.EncodeToString(*slice)
	return bson.MarshalValue(val)
}
