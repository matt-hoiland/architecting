package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserCredentials represents the document model stored in mongo.
type UserCredentials struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"        json:"_id,omitempty"`
	Email                string             `bson:"email"                json:"email"`
	Hash                 []byte             `bson:"hash"                 json:"-"`
	Salt                 []byte             `bson:"salt"                 json:"-"`
	Validated            bool               `bson:"validated"            json:"validated"`
	TimestampUpdatedLast time.Time          `bson:"ts_updated,omitempty" json:"ts_updated,omitempty"`
}
